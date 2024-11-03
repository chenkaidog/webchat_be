package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"time"
	"webchat_be/biz/dao"
	"webchat_be/biz/db/redis"
	"webchat_be/biz/model/consts"
	"webchat_be/biz/model/dto"
	"webchat_be/biz/model/errs"
	"webchat_be/biz/util/origin"
)

func ChatLimiter() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		sessId := ctx.Value(consts.SessionKeySessID).(string)
		redisClient := redis.GetRedisClient()

		// 每个会话1s只能请求一次接口
		qpsKey := fmt.Sprintf("chat_limiter_qps_%s", sessId)
		ok, err := redisClient.SetNX(ctx, qpsKey, true, time.Second).Result()
		if err == nil && !ok {
			dto.AbortWithErr(c, errs.TooManyRequest, http.StatusTooManyRequests)
			return
		}

		// 每个会话只能同时执行一次响应
		respondingKey := fmt.Sprintf("chat_limiter_responding_%s", sessId)
		ok, err = redisClient.SetNX(ctx, respondingKey, true, time.Second*5).Result()
		if err == nil && !ok {
			dto.AbortWithErr(c, errs.TooManyRequest, http.StatusTooManyRequests)
			return
		}

		defer redisClient.Del(ctx, respondingKey)

		c.Next(ctx)
	}
}

func LoginLimiter() app.HandlerFunc {
	// 相同IP在过去1h有过多登陆记录，
	return func(ctx context.Context, c *app.RequestContext) {
		loginRecordDao := dao.NewLoginRecordDao()
		originIp := origin.GetIp(c)
		hourAgo, _ := loginRecordDao.QueryByIP(ctx, originIp, time.Now().Add(-time.Hour))
		if len(hourAgo) > 20 {
			SetRiskIP(ctx, originIp)
			dto.AbortWithErr(c, errs.LoginReachLimit, http.StatusForbidden)
			return
		}

		dayAgo, _ := loginRecordDao.QueryByIP(ctx, originIp, time.Now().Add(-24*time.Hour))
		if len(dayAgo) > 40 {
			SetRiskIP(ctx, originIp)
			dto.AbortWithErr(c, errs.LoginReachLimit, http.StatusForbidden)
			return
		}

		c.Next(ctx)
	}
}

func parseRiskIpKey(ip string) string {
	return fmt.Sprintf("webchat_risk_ip:%s", ip)
}

func SetRiskIP(ctx context.Context, ip string) {
	riskIpKey := parseRiskIpKey(ip)
	value, err := redis.GetRedisClient().Incr(ctx, riskIpKey).Result()
	defer redis.GetRedisClient().ExpireAt(ctx, riskIpKey, time.Now().Add(24*time.Hour))
	if err == nil && value >= 3 {
		// IP触发3次风险操作，此时将封禁IP
		redis.GetRedisClient().HSet(ctx, "webchat_blocked_ip", ip, true)
	}
}

func IPIsAllowed() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if ipIsBlocked(ctx, origin.GetIp(c)) {
			dto.AbortWithErr(c, errs.RequestBlocked, http.StatusForbidden)
			return
		}

		c.Next(ctx)
	}
}

func ipIsBlocked(ctx context.Context, ip string) bool {
	blocked, err := redis.GetRedisClient().HGet(ctx, "webchat_blocked_ip", ip).Bool()
	if err != nil {
		return false
	}

	return blocked
}

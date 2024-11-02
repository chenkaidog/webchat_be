package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"time"
	redisclient "webchat_be/biz/db/redis"
)

// 缓存中添加登录状态
func appendLoginAccount(ctx context.Context, accountId, sessId string) error {
	limitKey := parseAccountLoginLimitKey(accountId)
	redisClient := redisclient.GetRedisClient()

	// 获取最近登录的会话记录
	results, err := redisClient.ZRevRangeWithScores(ctx, limitKey, 0, 0).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		hlog.CtxErrorf(ctx, "z range failed: %v", err)
		return err
	}

	validAccount := []redis.Z{
		{
			Score:  float64(time.Now().Unix()),
			Member: sessId,
		},
	}
	for _, result := range results {
		validAccount = append(validAccount, redis.Z{
			Score:  result.Score,
			Member: result.Member,
		})
	}

	pipeline := redisClient.Pipeline()
	pipeline.Del(ctx, limitKey)
	pipeline.ZAdd(ctx, limitKey, validAccount...)
	_, err = pipeline.Exec(ctx)
	return err
}

func parseAccountLoginLimitKey(accountId string) string {
	return fmt.Sprintf("account_login_limit_%s", accountId)
}

func RemoveSession(ctx context.Context, accountId, sessId string) error {
	return redisclient.GetRedisClient().ZRem(ctx, parseAccountLoginLimitKey(accountId), sessId).Err()
}

func SessionIsExpired(ctx context.Context, accountId, sessId string) bool {
	_, err := redisclient.GetRedisClient().
		ZScore(ctx, parseAccountLoginLimitKey(accountId), sessId).Result()

	// 查不到这个账户下对应的会话，判断为会话被挤占下线
	return errors.Is(err, redis.Nil)
}

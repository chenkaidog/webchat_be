package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"time"
	"webchat_be/biz/db/redis"
	"webchat_be/biz/model/dto"
	"webchat_be/biz/model/errs"
	"webchat_be/biz/util/origin"
)

func Limiter() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		key := fmt.Sprintf("qps_limiter_%s_%s", origin.GetIp(c), c.Path())
		ok, err := redis.GetRedisClient().SetNX(ctx, key, true, time.Second/10).Result()
		if err == nil && !ok {
			dto.AbortWithErr(c, errs.TooManyRequest, http.StatusTooManyRequests)
			return
		}

		c.Next(ctx)
	}
}

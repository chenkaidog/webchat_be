package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/rbcervilla/redisstore/v9"
	"net/http"
	"webchat_be/biz/db/redis"
	"webchat_be/biz/model/consts"
)

const (
	sessionStorePrefix = "webchat_session_"
)

func Session() app.HandlerFunc {
	store := NewRedisStore()
	store.Options(
		sessions.Options{
			Path:     "/",
			Domain:   "",
			MaxAge:   7 * 24 * 3600,
			Secure:   false, // https not ready
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		},
	)
	return sessions.Many(
		[]string{
			consts.SessionNameAccount,        // 正常用户会话
			consts.SessionNameRegisterVerify, // 注册会话
			consts.SessionNameResetVerify,    // 重置密码会话
		}, store)
}

type RedisStore struct {
	*redisstore.RedisStore
}

func (r *RedisStore) Options(opts sessions.Options) {
	r.RedisStore.Options(*opts.ToGorillaOptions())
}

func NewRedisStore() *RedisStore {
	redisStore, err := redisstore.NewRedisStore(context.Background(), redis.GetRedisClient())
	if err != nil {
		panic(err)
	}
	redisStore.KeyPrefix(sessionStorePrefix)
	return &RedisStore{
		RedisStore: redisStore,
	}
}

package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"net/http"
	"webchat_be/biz/model/consts"
	"webchat_be/biz/util/origin"
)

func LoginStateVerify() []app.HandlerFunc {
	return []app.HandlerFunc{
		LoginSession(),
		CSRF(),
	}
}

func LoginSession() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.DefaultMany(c, consts.SessionNameAccount)
		originIP, ok := session.Get(consts.SessionKeyLoginIP).(string)
		if ok && originIP != origin.GetIp(c) {
			c.AbortWithMsg("user not login", http.StatusUnauthorized)
			return
		}
		originDevice, ok := session.Get(consts.SessionKeyDevice).(string)
		if ok && originDevice != origin.GetDevice(c) {
			c.AbortWithMsg("user not login", http.StatusUnauthorized)
			return
		}

		accountId, ok := session.Get(consts.SessionKeyAccountId).(string)
		if !ok {
			c.AbortWithMsg("user not login", http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, consts.SessionKeyLoginIP, originIP)
		ctx = context.WithValue(ctx, consts.SessionKeyDevice, originDevice)
		ctx = context.WithValue(ctx, consts.SessionKeyAccountId, accountId)

		c.Next(ctx)
	}
}

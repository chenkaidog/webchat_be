package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"net/http"
	"webchat_be/biz/handler/service"
	"webchat_be/biz/model/consts"
	"webchat_be/biz/model/dto"
	"webchat_be/biz/model/errs"
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
		originIP, ok1 := session.Get(consts.SessionKeyLoginIP).(string)
		originDevice, ok2 := session.Get(consts.SessionKeyDevice).(string)
		accountId, ok3 := session.Get(consts.SessionKeyAccountId).(string)
		sessId, ok4 := session.Get(consts.SessionKeySessID).(string)
		if ok1 && originIP != origin.GetIp(c) &&
			ok2 && originDevice != origin.GetDevice(c) ||
			!ok3 || !ok4 {
			dto.AbortWithErr(c, errs.Unauthorized, http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, consts.SessionKeySessID, sessId)
		ctx = context.WithValue(ctx, consts.SessionKeyLoginIP, originIP)
		ctx = context.WithValue(ctx, consts.SessionKeyDevice, originDevice)
		ctx = context.WithValue(ctx, consts.SessionKeyAccountId, accountId)

		if service.SessionIsExpired(ctx, accountId, sessId) {
			dto.AbortWithErr(c, errs.Unauthorized, http.StatusUnauthorized)
			return
		}

		c.Next(ctx)
	}
}

package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"net/http"
	"webchat_be/biz/model/consts"
	"webchat_be/biz/model/dto"
	"webchat_be/biz/model/errs"
	"webchat_be/biz/util/encode"
)

func CSRF() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.DefaultMany(c, consts.SessionNameAccount)
		salt, _ := session.Get(consts.SessionKeyCsrfSalt).(string)
		encodedToken, _ := session.Get(consts.SessionKeyCsrfToken).(string)

		csrfToken := c.Request.Header.Get(consts.HeaderKeyCsrfToken)
		if encode.EncodePassword(salt, csrfToken) != encodedToken {
			dto.AbortWithErr(c, errs.CsrfTokenInvalid, http.StatusUnauthorized)
			return
		}

		c.Next(ctx)
	}
}

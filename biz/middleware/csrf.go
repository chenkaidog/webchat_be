package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"net/http"
	"webchat_be/biz/model/consts"
	"webchat_be/biz/util/encode"
)

func CSRF() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.DefaultMany(c, consts.SessionNameAccount)
		salt, ok := session.Get(consts.SessionKeyCsrfSalt).(string)
		if !ok {
			c.AbortWithMsg("csrf token invalid", http.StatusUnauthorized)
			return
		}
		encodedToken, ok := session.Get(consts.SessionKeyCsrfToken).(string)
		if !ok {
			c.AbortWithMsg("csrf token invalid", http.StatusUnauthorized)
			return
		}

		csrfToken := c.Request.Header.Get(consts.HeaderKeyCsrfToken)
		if encode.EncodePassword(salt, csrfToken) != encodedToken {
			c.AbortWithMsg("csrf token invalid", http.StatusUnauthorized)
			return
		}

		c.Next(ctx)
	}
}

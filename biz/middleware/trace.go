package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"webchat_be/biz/util/id_gen"
	"webchat_be/biz/util/trace_info"
)

const (
	headerKeyTraceId = "X-Trace-ID"
	headerKeyLogId   = "X-Log-ID"
	headerKeySpanId  = "X-Span-ID"
)

func TraceContext() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		logID := c.Request.Header.Get(headerKeyLogId)
		if logID == "" {
			logID = id_gen.NewLogID()
		}

		ctx = trace_info.WithTrace(
			ctx,
			trace_info.TraceInfo{
				LogID: logID,
			})

		c.Next(ctx)

		c.Header(headerKeyLogId, logID)
	}
}

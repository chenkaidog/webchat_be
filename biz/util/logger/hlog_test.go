package logger

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"testing"
)

func TestHlog(t *testing.T) {
	Init()

	ctx := context.Background()
	hlog.CtxInfof(ctx, "test data: %d, %s", 123, "ttt")
}

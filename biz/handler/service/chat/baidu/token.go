package baidu

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"net/http"
	"runtime/debug"
	"time"
	"webchat_be/biz/config"
	"webchat_be/biz/util/http_client"
	"webchat_be/biz/util/id_gen"
	"webchat_be/biz/util/trace_info"
)

func Init() {
	ticker := time.NewTicker(time.Second)

	go func() {
		ctx := trace_info.WithTrace(
			context.Background(),
			trace_info.TraceInfo{
				LogID: id_gen.NewLogID(),
			})

		defer func() {
			if rec := recover(); rec != nil {
				hlog.CtxErrorf(ctx, "panic: %v\n%s", rec, debug.Stack())
			}
			ticker.Stop()
		}()

		for range ticker.C {
			accessInfo, err := refreshAccessToken(ctx, config.GetBaiduConf().AppKey, config.GetBaiduConf().AppSecret)
			if err != nil {
				ticker.Reset(time.Second)
				continue
			}
			if accessInfo != nil {
				localAccessToken = accessInfo
				expireIn := localAccessToken.ExpiresIn * 99 / 100
				ticker.Reset(time.Duration(expireIn) * time.Second)
			}
		}
	}()
}

var localAccessToken *AppAccessInfo

func refreshAccessToken(ctx context.Context, appKey, appSecret string) (*AppAccessInfo, error) {
	req, err := http.NewRequest(http.MethodGet, accessUrl, nil)
	if err != nil {
		panic(err)
	}

	param := req.URL.Query()
	param.Set("client_id", appKey)
	param.Set("client_secret", appSecret)
	req.URL.RawQuery = param.Encode()

	resp, err := http_client.NewHttpClient(false).Do(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "http request err: %v", err)
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		hlog.CtxErrorf(ctx, "read resp body err: %v", err)
		return nil, err
	}

	var accessInfo *AppAccessInfo
	if err = json.Unmarshal(data, &accessInfo); err != nil {
		hlog.CtxErrorf(ctx, "unmarshal err: %v", err)
		return nil, err
	}

	return accessInfo, nil
}

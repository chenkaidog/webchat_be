package http_client

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"net/url"
	"webchat_be/biz/config"
)

func NewHttpClient(proxyOn bool) *http.Client {
	if proxyOn {
		proxyUrl, err := url.Parse(config.GetHttpProxyConf().Url)
		if err != nil {
			hlog.Errorf("http proxy url parse error: %s", err.Error())
		}

		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}

	return http.DefaultClient
}

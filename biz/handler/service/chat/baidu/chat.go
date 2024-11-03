package baidu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"net/http"
	"strings"
	"webchat_be/biz/model/domain"
	"webchat_be/biz/model/errs"
	"webchat_be/biz/util/http_client"
	"webchat_be/biz/util/sse_client"
)

type Chat struct {
	model string
}

func NewChat(model string) *Chat {
	return &Chat{
		model: model,
	}
}

func (c *Chat) StreamChat(ctx context.Context, contents []*domain.ChatContent) (chan *domain.StreamingResp, errs.Error) {
	httpReq, err := c.newStreamChatRequest(ctx, contents)
	if err != nil {
		return nil, errs.ServerError
	}
	httpResp, err := http_client.NewHttpClient(false).Do(httpReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "http request err: %v", err)
		return nil, errs.ServerError
	}
	if bizErr := c.handleRespStatus(ctx, httpResp); bizErr != nil {
		return nil, bizErr
	}

	return sse_client.HandleSeeResp(ctx, httpResp, parseStreamingResp), nil
}

func (c *Chat) newStreamChatRequest(ctx context.Context, contents []*domain.ChatContent) (*http.Request, error) {
	var messages []*Message
	for _, content := range contents {
		messages = append(messages, &Message{
			Role:    roleMapper[content.Role],
			Content: content.Content,
		},
		)
	}

	reqBody, err := json.Marshal(&ChatCreateReq{
		Stream:   true,
		Messages: messages,
	},
	)
	if err != nil {
		hlog.CtxErrorf(ctx, "json marshal err: %v", err)
		return nil, err
	}

	requestUrl := fmt.Sprintf(chatUrl, c.model)
	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		hlog.CtxErrorf(ctx, "new request err: %v", err)
		return nil, err
	}

	if localAccessToken == nil {
		hlog.CtxErrorf(ctx, "access token not init: %v", err)
		return nil, fmt.Errorf("access token not init")
	}

	param := req.URL.Query()
	param.Set("access_token", localAccessToken.AccessToken)
	req.URL.RawQuery = param.Encode()

	hlog.CtxDebugf(ctx, "request: %+v", req)

	return req, nil
}

func (c *Chat) handleRespStatus(ctx context.Context, httpResp *http.Response) errs.Error {
	if httpResp.StatusCode != http.StatusOK {
		respContent, _ := io.ReadAll(httpResp.Body)
		hlog.CtxErrorf(ctx, "status_code: %d, error_msg: %s", httpResp.StatusCode, respContent)
		return errs.ServerError
	} else if strings.Contains(httpResp.Header.Get("Content-Type"), "json") {
		respContent, _ := io.ReadAll(httpResp.Body)
		var errRespBody ChatError
		if err := json.Unmarshal(respContent, &errRespBody); err != nil {
			hlog.CtxErrorf(ctx, "json unmarshal err: %v", err)
			return errs.ServerError
		}

		// https://cloud.baidu.com/doc/WENXINWORKSHOP/s/tlmyncueh
		switch errRespBody.ErrorCode {
		case 4, 18:
			hlog.CtxErrorf(ctx, "request limit reached err: %v", errRespBody.ErrorMsg)
			return errs.ChatRateLimitReached
		case 17:
			hlog.CtxWarnf(ctx, "exceed quote limit: %s", errRespBody.ErrorMsg)
			return errs.ExceedQuoteLimit
		default:
			hlog.CtxErrorf(ctx, "api err, %d:%v", errRespBody.ErrorCode, errRespBody.ErrorMsg)
			return errs.ServerError
		}
	}

	return nil
}

func parseStreamingResp(ctx context.Context, data []byte) *domain.StreamingResp {
	if data == nil {
		return nil
	}

	var respBody ChatCreateResp
	if err := json.Unmarshal(data, &respBody); err != nil {
		hlog.CtxErrorf(ctx, "json unmarshal err: %v", err)
		return nil
	}

	if respBody.Error != "" {
		hlog.CtxErrorf(ctx, "request err: %s", respBody.ErrorDescription)
		return &domain.StreamingResp{
			Msg:     "baidu platform error",
			IsEnd:   true,
			IsError: true,
		}
	}

	return &domain.StreamingResp{
		Msg:     respBody.Result,
		IsEnd:   respBody.IsEnd,
		IsError: false,
	}
}

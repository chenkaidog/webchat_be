package baidu

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"net/http"
	"webchat_be/biz/model/domain"
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

func (c *Chat) StreamChat(ctx context.Context, contents []*domain.ChatContent) (chan *domain.StreamingResp, error) {
	httpReq, err := c.newStreamChatRequest(ctx, contents)
	if err != nil {
		return nil, err
	}
	httpResp, err := http_client.NewHttpClient(false).Do(httpReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "http request err: %v", err)
		return nil, err
	}
	if httpResp.StatusCode != http.StatusOK {
		respContent, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}

		hlog.CtxErrorf(ctx, "status_code: %d, error_msg: %s", httpResp.StatusCode, respContent)
		return nil, errors.New(string(respContent))
	}

	streamChan := sse_client.HandleSeeResp(ctx, httpResp, parseStreamingResp)

	return streamChan, nil
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

	return req, nil
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

	if respBody.ErrorCode != 0 || respBody.Error != "" {
		hlog.CtxErrorf(ctx, "request err: %s. %s.", respBody.ErrorMsg, respBody.ErrorDescription)
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

package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"net/http"
	"webchat_be/biz/config"
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

	httpResp, err := http_client.NewHttpClient(true).Do(httpReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "http request failed: %v", err)
		return nil, err
	}
	if httpResp.StatusCode != http.StatusOK {
		respContent, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}

		hlog.CtxErrorf(ctx, "status_code: %d, error_msg: %s", httpResp.StatusCode, respContent)
		return nil, errors.New("request fails")
	}

	return sse_client.HandleSeeResp(ctx, httpResp, parseStreamingResp), nil
}

func (c *Chat) newStreamChatRequest(ctx context.Context, contents []*domain.ChatContent) (*http.Request, error) {
	var messages []Message
	for _, content := range contents {
		messages = append(
			messages,
			Message{
				Role:    roleMapper[content.Role],
				Content: content.Content,
			},
		)
	}

	reqBody, err := json.Marshal(
		&ChatCreateReq{
			Model:    c.model,
			Stream:   true,
			Messages: messages,
		},
	)
	if err != nil {
		hlog.CtxErrorf(ctx, "json marshal err: %v", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, chatUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		hlog.CtxErrorf(ctx, "new http request err: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GetOpenAIConf().ApiKey)
	hlog.CtxDebugf(ctx, "request: %+v", req)

	return req, nil
}

func parseStreamingResp(ctx context.Context, data []byte) *domain.StreamingResp {
	if data == nil || bytes.Equal(data, []byte("[DONE]")) {
		return nil
	}

	var respBody ChatCreateResp
	if err := json.Unmarshal(data, &respBody); err != nil {
		hlog.CtxErrorf(ctx, "json unmarshal err: %v", err)
		return nil
	}

	if len(respBody.Choices) > 0 {
		choice := respBody.Choices[0]
		if choice.FinishReason == finishReasonStop {
			return &domain.StreamingResp{
				Msg:     "",
				IsEnd:   true,
				IsError: false,
			}
		}

		return &domain.StreamingResp{
			Msg:     choice.Delta.Content,
			IsEnd:   false,
			IsError: false,
		}
	}

	return nil
}

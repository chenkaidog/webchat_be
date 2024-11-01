package chat

import (
	"context"
	"webchat_be/biz/chat/baidu"
	"webchat_be/biz/chat/openai"
	"webchat_be/biz/model/domain"
)

type Inf interface {
	StreamChat(context.Context, []*domain.ChatContent) (chan *domain.StreamingResp, error)
}

func NewChatImpl(platform, model string) Inf {
	switch platform {
	case domain.PlatformBaidu:
		return baidu.NewChat(model)
	case domain.PlatformOpenai:
		return openai.NewChat(model)
	}

	return nil
}

package openai

import "webchat_be/biz/model/domain"

var roleMapper = map[domain.Role]string{
	domain.RoleSystem:    roleSystem,
	domain.RoleUser:      roleUser,
	domain.RoleAssistant: roleAssistant,
}

const chatUrl = "https://api.openai.com/v1/chat/completions"

const (
	roleAssistant = "assistant"
	roleUser      = "user"
	roleSystem    = "system"
)

type ChatCreateReq struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCreateResp struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []*Choice `json:"choices"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Delta        Delta   `json:"delta"`
}

type Delta struct {
	Content string `json:"content"`
}

const finishReasonStop = "stop"

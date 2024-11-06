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

// https://platform.openai.com/docs/api-reference/chat/create

type ChatCreateReq struct {
	Model               string    `json:"model"`
	Stream              bool      `json:"stream"`
	MaxCompletionTokens int       `json:"max_completion_tokens"`
	User                string    `json:"user"`
	Messages            []Message `json:"messages"`
}

type StreamOptions struct {
	IncludeUsage bool `json:"include_usage"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Content struct {
	Type    string `json:"type"`
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

type ErrorResp struct {
	Error struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    string      `json:"code"`
	} `json:"error"`
}

const finishReasonStop = "stop"

package domain

type ChatContent struct {
	Role    Role
	Content string
}

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

const (
	PlatformBaidu  = "baidu"
	PlatformOpenai = "openai"
)

type StreamingResp struct {
	Msg     string
	IsEnd   bool
	IsError bool
}

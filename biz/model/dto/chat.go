package dto

type Message struct {
	Role    Role   `form:"role" json:"role" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type ChatCreateReq struct {
	ModelId  string    `json:"model_id" binding:"required"`
	Messages []Message `json:"messages" binding:"required,max=20"`
}

type ChatCreateResp struct {
}

type ChatStreamResp struct {
	CreatedAt int64  `json:"created_at,omitempty"`
	IsEnd     bool   `json:"is_end,omitempty"`
	Content   string `json:"content,omitempty"`
}

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

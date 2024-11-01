package handler

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sse"
	"net/http"
	"time"
	"webchat_be/biz/chat"
	"webchat_be/biz/dao"
	"webchat_be/biz/model/domain"
	"webchat_be/biz/model/dto"
)

// StreamingChat 用户对话接口
//
//	@Tags		chat
//	@Accept		application/json
//	@Produce	text/event-stream
//	@Param		req				body		dto.ChatCreateReq	true	"request body"
//	@Param		X-CSRF-TOKEN	header		string				true	"csrf token"
//	@Success	200				{object}	dto.ChatStreamResp
//	@Router		/api/v1/chat/stream [POST]
func StreamingChat(ctx context.Context, c *app.RequestContext) {
	var chatCreateReq dto.ChatCreateReq
	if stdErr := c.BindAndValidate(&chatCreateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		c.AbortWithMsg("request body invalid", http.StatusBadRequest)
		return
	}

	model, err := dao.NewModelDao().QueryByModelId(ctx, chatCreateReq.ModelId)
	if err != nil {
		hlog.CtxInfof(ctx, "QueryByModelId fail, %v", err)
		c.AbortWithMsg("internal error", http.StatusInternalServerError)
		return
	}
	if model == nil {
		c.AbortWithMsg("model not found", http.StatusBadRequest)
		return
	}

	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	streamChan, err := chat.NewChatImpl(model.Platform, model.Name).
		StreamChat(cancelCtx, parseStreamChatReq(&chatCreateReq))
	if err != nil {
		hlog.CtxInfof(ctx, "StreamChat error, %v", err)
		c.AbortWithMsg("internal error", http.StatusInternalServerError)
		return
	}

	c.SetStatusCode(http.StatusOK)
	ssePublisher := sse.NewStream(c)
	for chatResp := range streamChan {
		if chatResp.IsEnd {
			if chatResp.IsError {
				errorPublish(ssePublisher, chatResp.Msg)
			} else {
				endPublish(ssePublisher)
			}

			return
		}

		if err := msgPublish(ssePublisher, chatResp.Msg); err != nil {
			return
		}
	}
}

var roleMapper = map[dto.Role]domain.Role{
	dto.RoleSystem:    domain.RoleSystem,
	dto.RoleUser:      domain.RoleUser,
	dto.RoleAssistant: domain.RoleAssistant,
}

func parseStreamChatReq(chatCreateReq *dto.ChatCreateReq) []*domain.ChatContent {
	var result []*domain.ChatContent

	for _, msg := range chatCreateReq.Messages {
		result = append(result, &domain.ChatContent{
			Role:    roleMapper[msg.Role],
			Content: msg.Content,
		})
	}

	return result
}

func errorPublish(stream *sse.Stream, msg string) {
	resp := &dto.ChatStreamResp{
		CreatedAt: time.Now().Unix(),
		IsEnd:     true,
		Content:   msg,
	}
	data, _ := json.Marshal(resp)
	_ = stream.Publish(&sse.Event{
		Event: "error",
		Data:  data,
	})
}

func endPublish(stream *sse.Stream) {
	resp := &dto.ChatStreamResp{
		CreatedAt: time.Now().Unix(),
		IsEnd:     true,
		Content:   "",
	}
	data, _ := json.Marshal(resp)
	_ = stream.Publish(&sse.Event{
		Event: "data",
		Data:  data,
	})
}

func msgPublish(stream *sse.Stream, content string) error {
	resp := &dto.ChatStreamResp{
		CreatedAt: time.Now().Unix(),
		IsEnd:     false,
		Content:   content,
	}
	data, _ := json.Marshal(resp)

	return stream.Publish(&sse.Event{
		Event: "data",
		Data:  data,
	})
}

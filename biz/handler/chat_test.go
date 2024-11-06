package handler

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"webchat_be/biz/model/dto"
)

func Test_parseStreamChatReq(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var message []dto.Message
		for i := 0; i < 21; i++ {
			message = append(message, dto.Message{})
		}
		result := parseStreamChatReq(&dto.ChatCreateReq{
			Messages: message,
		})
		assert.Equal(t, len(result), 11)
	})

	t.Run("", func(t *testing.T) {
		var message []dto.Message
		for i := 0; i < 5; i++ {
			message = append(message, dto.Message{})
		}
		result := parseStreamChatReq(&dto.ChatCreateReq{
			Messages: message,
		})
		assert.Equal(t, len(result), 5)
	})

	t.Run("", func(t *testing.T) {
		var message []dto.Message
		for i := 0; i < 11; i++ {
			message = append(message, dto.Message{})
		}
		result := parseStreamChatReq(&dto.ChatCreateReq{
			Messages: message,
		})
		assert.Equal(t, len(result), 11)
	})
}

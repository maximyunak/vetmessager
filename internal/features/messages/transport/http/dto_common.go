package messages_transport_http

import (
	"time"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type MessageDTOResponse struct {
	ID               int        `json:"id"`
	ChatID           int        `json:"chat_id"`
	SenderID         int        `json:"sender_id"`
	ReplyToMessageID *int       `json:"reply_to_message_id"`
	Content          *string    `json:"content"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

func MessageDTOFromDomain(message domain.Message) MessageDTOResponse {
	return MessageDTOResponse{
		ID:               message.ID,
		ChatID:           message.ChatID,
		SenderID:         message.SenderID,
		ReplyToMessageID: message.ReplyToMessageID,
		Content:          message.Content,
		CreatedAt:        message.CreatedAt,
		UpdatedAt:        message.UpdatedAt,
		DeletedAt:        message.DeletedAt,
	}
}

func MessagesDTOFromDomain(messages []domain.Message) []MessageDTOResponse {
	messagesDTO := make([]MessageDTOResponse, len(messages))

	for i, message := range messages {
		messagesDTO[i] = MessageDTOFromDomain(message)
	}

	return messagesDTO
}

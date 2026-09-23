package domain

import (
	"fmt"
	"time"

	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

type Message struct {
	ID               int        `json:"id" db:"id"`
	ChatID           int        `json:"chat_id" db:"chat_id"`
	SenderID         int        `json:"sender_id" db:"sender_id"`
	ReplyToMessageID *int       `json:"reply_to_message_id" db:"reply_to_message_id"`
	Content          *string    `json:"content" db:"content"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at" db:"deleted_at"`
}

func NewMessageUninitialized(chatId int, senderId int, replyId *int, content *string) Message {
	return Message{
		ID:               UninitializedID,
		ChatID:           chatId,
		SenderID:         senderId,
		ReplyToMessageID: replyId,
		Content:          content,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (m Message) Validate() error {
	if m.ChatID <= 0 {
		return fmt.Errorf(
			"invalid `chat_id`: %d: %w",
			m.ChatID,
			core_errors.ErrInvalidArgument,
		)
	}

	if m.SenderID <= 0 {
		return fmt.Errorf(
			"invalid `sender_id`: %d: %w",
			m.SenderID,
			core_errors.ErrInvalidArgument,
		)
	}

	contentLength := len([]rune(*m.Content))
	if contentLength < 1 || contentLength > 5000 {
		return fmt.Errorf(
			"invalid `content` length: %d: %w",
			contentLength,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

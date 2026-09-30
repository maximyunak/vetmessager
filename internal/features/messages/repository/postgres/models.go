package messages_postgres_repository

import (
	"time"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type MessageModel struct {
	ID               int        `db:"id"`
	ChatID           int        `db:"chat_id"`
	SenderID         int        `db:"sender_id"`
	ReplyToMessageID *int       `db:"reply_to_message_id"`
	Content          *string    `db:"content"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
	DeletedAt        *time.Time `db:"deleted_at"`
}

func messageDomainFromModel(model MessageModel) domain.Message {
	return domain.Message{
		ID:               model.ID,
		ChatID:           model.ChatID,
		SenderID:         model.SenderID,
		ReplyToMessageID: model.ReplyToMessageID,
		Content:          model.Content,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		DeletedAt:        model.DeletedAt,
	}
}

func messageDomainsFromModels(messages []MessageModel) []domain.Message {
	domains := make([]domain.Message, len(messages))

	for i, m := range messages {
		domains[i] = messageDomainFromModel(m)
	}

	return domains
}

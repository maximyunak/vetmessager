package chat_postgres_repository

import (
	"time"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type ChatModel struct {
	ID        int
	Name      *string
	Type      string
	CreatorID *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func chatDomainsFromModels(chats []ChatModel) []domain.Chat {
	domains := make([]domain.Chat, len(chats))

	for i, chat := range chats {
		domains[i] = domain.Chat{
			ID:        chat.ID,
			Name:      chat.Name,
			Type:      chat.Type,
			CreatorID: chat.CreatorID,
			CreatedAt: chat.CreatedAt,
			UpdatedAt: chat.UpdatedAt,
		}
	}

	return domains
}

func chatDomainFromModel(chat ChatModel) domain.Chat {
	return domain.Chat{
		ID:        chat.ID,
		Name:      chat.Name,
		Type:      chat.Type,
		CreatorID: chat.CreatorID,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}

type ChatMemberModel struct {
	ChatID    int       `db:"chat_id"`
	UserID    int       `db:"user_id"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func chatMemberDomainFromModel(model ChatMemberModel) domain.ChatMember {
	return domain.ChatMember{
		ChatID:    model.ChatID,
		UserID:    model.UserID,
		Role:      model.Role,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

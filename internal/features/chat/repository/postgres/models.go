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

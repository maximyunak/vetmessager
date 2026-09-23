package messages_service

import (
	"context"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type MessageService struct {
	MessageRepository MessageRepository
	TokenManager      auth.TokenManager
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error)
}

func NewMessageService(messageRepository MessageRepository, tokenManager auth.TokenManager) *MessageService {
	return &MessageService{MessageRepository: messageRepository, TokenManager: tokenManager}
}

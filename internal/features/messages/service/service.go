package messages_service

import (
	"context"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	"github.com/maximyunak/vetmessager/internal/core/domain"
	"github.com/maximyunak/vetmessager/internal/realtime/websocket"
)

type MessageService struct {
	MessageRepository MessageRepository
	TokenManager      auth.TokenManager
	hub               *websocket.Hub
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error)
}

func NewMessageService(messageRepository MessageRepository, tokenManager auth.TokenManager, hub *websocket.Hub) *MessageService {
	return &MessageService{MessageRepository: messageRepository, TokenManager: tokenManager, hub: hub}
}

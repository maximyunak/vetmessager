package messages_service

import (
	"context"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_websocket "github.com/maximyunak/vetmessager/internal/core/transport/websocket"
)

type MessageService struct {
	MessageRepository MessageRepository
	TokenManager      auth.TokenManager
	hub               *core_websocket.Hub
	ChatRepository    ChatRepository
}

type ChatRepository interface {
	GetChatMember(
		ctx context.Context,
		chatID int,
		userID int,
	) (domain.ChatMember, error)
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error)
	GetMessage(
		ctx context.Context,
		messageID int,
	) (domain.Message, error)
	DeleteMessage(
		ctx context.Context,
		messageID int,
	) error
}

func NewMessageService(messageRepository MessageRepository, tokenManager auth.TokenManager, hub *core_websocket.Hub, ChatRepository ChatRepository) *MessageService {
	return &MessageService{MessageRepository: messageRepository, TokenManager: tokenManager, hub: hub, ChatRepository: ChatRepository}
}

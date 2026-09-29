package messages_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_websocket "github.com/maximyunak/vetmessager/internal/core/transport/websocket"
)

func (s *MessageService) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {
	if err := message.Validate(); err != nil {
		return domain.Message{}, err
	}

	// check exists user in chat
	member, err := s.ChatRepository.GetChatMember(ctx, message.ChatID, message.SenderID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("check chat membership: %w", err)
	}

	// check user mute
	if member.IsMuted {
		return domain.Message{}, fmt.Errorf("member is muted: %w", core_errors.ErrForbidden)
	}

	// create message
	message, err = s.MessageRepository.CreateMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	// send to ws
	event := core_websocket.NewMessageEvent("message.created", message)

	s.hub.BroadcastToChat(message.ChatID, event)

	return message, nil
}

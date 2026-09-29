package messages_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_websocket "github.com/maximyunak/vetmessager/internal/core/transport/websocket"
)

func (s *MessageService) EditMessage(ctx context.Context, patch domain.MessagePatch) (domain.Message, error) {

	// check exists user in chat
	member, err := s.ChatRepository.GetChatMember(ctx, patch.ChatID, patch.SenderID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("check chat membership: %w", err)
	}

	// check user mute
	if member.IsMuted {
		return domain.Message{}, fmt.Errorf("member is muted: %w", core_errors.ErrForbidden)
	}

	// check message
	message, err := s.MessageRepository.GetMessage(ctx, patch.ID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("get message: %w", err)
	}

	// check message status
	if message.DeletedAt != nil {
		return domain.Message{}, core_errors.ErrNotFound
	}

	// check message owner
	if message.SenderID != patch.SenderID {
		return domain.Message{}, core_errors.ErrForbidden
	}

	// check message chat = chatId from request
	if message.ChatID != patch.ChatID {
		return domain.Message{}, core_errors.ErrNotFound
	}

	// apply patch to message
	if err := message.ApplyPatch(patch); err != nil {
		return domain.Message{}, fmt.Errorf("apply patch: %w", err)
	}

	// edit message
	message, err = s.MessageRepository.EditMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("edit message: %w", err)
	}

	// send to ws
	event := core_websocket.NewMessageEvent("message.edited", message)

	s.hub.BroadcastToChat(message.ChatID, event)

	return message, nil
}

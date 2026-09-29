package messages_service

import (
	"context"
	"fmt"

	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_websocket "github.com/maximyunak/vetmessager/internal/core/transport/websocket"
)

func (s *MessageService) DeleteMessage(ctx context.Context, userID int, chatID int, messageID int) error {
	// check exists user in chat
	member, err := s.ChatRepository.GetChatMember(ctx, chatID, userID)
	if err != nil {
		return fmt.Errorf("check chat membership: %w", err)
	}

	// get message
	message, err := s.MessageRepository.GetMessage(ctx, messageID)
	if err != nil {
		return fmt.Errorf("get message: %w", err)
	}

	// check message status
	if message.DeletedAt != nil {
		return core_errors.ErrNotFound
	}

	// check user role
	isAdmin := member.Role == "admin"
	isOwner := member.Role == "owner"
	isAuthor := message.SenderID == userID

	if !isAdmin && !isOwner && !isAuthor {
		return core_errors.ErrForbidden
	}

	// check message chat = chatId from request
	if message.ChatID != chatID {
		return core_errors.ErrNotFound
	}

	// delete message
	err = s.MessageRepository.DeleteMessage(ctx, messageID)
	if err != nil {
		return fmt.Errorf("delete message: %w", err)
	}

	// send to ws
	event := core_websocket.NewMessageEvent("message.deleted", message)

	s.hub.BroadcastToChat(message.ChatID, event)

	return nil
}

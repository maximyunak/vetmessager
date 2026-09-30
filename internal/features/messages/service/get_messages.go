package messages_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

func (s *MessageService) GetMessages(ctx context.Context, userID int, chatID int, limit int, offset int) ([]domain.Message, int, error) {

	// Check that user is a member of the chat
	_, err := s.ChatRepository.GetChatMember(ctx, chatID, userID)
	if err != nil {
		return nil, 0, fmt.Errorf("check chat membership: %w", err)
	}

	// get chat messages
	messages, err := s.MessageRepository.GetMessages(ctx, chatID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get messages: %w", err)
	}
	total, err := s.MessageRepository.GetMessagesCount(ctx, chatID)
	if err != nil {
		return nil, 0, fmt.Errorf("get messages count: %w", err)
	}
	return messages, total, nil
}

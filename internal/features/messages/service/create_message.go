package messages_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

func (s *MessageService) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {
	if err := message.Validate(); err != nil {
		return domain.Message{}, err
	}

	message, err := s.MessageRepository.CreateMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create user: %w", err)
	}

	return message, nil
}

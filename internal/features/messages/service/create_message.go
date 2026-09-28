package messages_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	"github.com/maximyunak/vetmessager/internal/realtime/websocket"
)

func (s *MessageService) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {
	if err := message.Validate(); err != nil {
		return domain.Message{}, err
	}

	message, err := s.MessageRepository.CreateMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create user: %w", err)
	}

	// send to ws
	event := &websocket.Event{
		Type: "message.created",
		Payload: websocket.Payload{
			ChatID: message.ChatID,
			UserID: message.SenderID,
			Message: &websocket.Message{
				ID:      message.ID,
				Content: message.Content,
			},
		},
	}

	s.hub.BroadcastToChat(message.ChatID, event)

	return message, nil
}

package core_websocket

import "github.com/maximyunak/vetmessager/internal/core/domain"

func NewMessageEvent(messageType string, message domain.Message) *Event {
	return &Event{
		Type: messageType,
		Payload: Payload{
			ChatID: message.ChatID,
			UserID: message.SenderID,
			Message: &Message{
				ID:      message.ID,
				Content: message.Content,
			},
		},
	}
}

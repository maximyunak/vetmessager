package messages_service

type MessageService struct {
	MessageRepository MessageRepository
}

type MessageRepository interface {
	CreateMessage()
}

func NewMessageService(messageRepository MessageRepository) *MessageService {
	return &MessageService{MessageRepository: messageRepository}
}

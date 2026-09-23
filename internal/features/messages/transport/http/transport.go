package messages_transport_http

import (
	"context"
	"net/http"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_http_server "github.com/maximyunak/vetmessager/internal/core/transport/http/server"
)

type MessagesHTTPHandler struct {
	messagesService MessagesService
}

type MessagesService interface {
	CreateMessage(
		ctx context.Context, messageDomain domain.Message,
	) (domain.Message, error)
}

func NewMessagesHTTPHandler(messagesService MessagesService) *MessagesHTTPHandler {
	return &MessagesHTTPHandler{
		messagesService,
	}
}

func (h *MessagesHTTPHandler) PublicRoutes() []core_http_server.Route {
	return []core_http_server.Route{}
}

func (h *MessagesHTTPHandler) ProtectedRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/chat/{chatId}",
			Handler: h.CreateMessage,
		},
	}
}

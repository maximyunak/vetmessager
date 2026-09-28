package core_websocket

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/maximyunak/vetmessager/internal/core/auth"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
	core_http_server "github.com/maximyunak/vetmessager/internal/core/transport/http/server"
	chat_postgres_repository "github.com/maximyunak/vetmessager/internal/features/chat/repository/postgres"
)

type MessageWsHandler struct {
	hub            *Hub
	chatRepository chatRepository
}

type chatRepository interface {
	GetUserChatsIDs(ctx context.Context, userID int) ([]int, error)
}

func NewHandler(h *Hub, chatRepository *chat_postgres_repository.ChatRepository) *MessageWsHandler {
	return &MessageWsHandler{
		hub:            h,
		chatRepository: chatRepository,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *MessageWsHandler) Connect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke call ws handler")

	userID, ok := auth.UserIDFromContext(ctx)

	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Unauthorized")
		return
	}
	chatIDs, err := h.chatRepository.GetUserChatsIDs(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "get user chats IDs")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	chats := make(map[int]struct{})

	for _, chatID := range chatIDs {
		chats[chatID] = struct{}{}
	}

	client := &Client{
		Connection: conn,
		Message:    make(chan *Event, 10),
		ID:         userID,
		Chats:      chats,
	}

	h.hub.register <- client

	go client.WriteMessage()
	client.ReadMessage(h.hub)
}

func (h *MessageWsHandler) ProtectedRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/ws",
			Handler: h.Connect,
		},
	}
}

package messages_transport_websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/maximyunak/vetmessager/internal/core/auth"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
	core_http_server "github.com/maximyunak/vetmessager/internal/core/transport/http/server"
)

type MessageWsHandler struct {
	hub *Hub
}

func NewHandler(h *Hub) *MessageWsHandler {
	return &MessageWsHandler{
		hub: h,
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

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		responseHandler.ErrorResponse(err, "upgrade error")
		return
	}

	client := &Client{
		Connection: conn,
		Message:    make(chan *Event, 10),
		ID:         userID,
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

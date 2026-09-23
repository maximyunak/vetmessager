package messages_transport_http

import (
	"net/http"
	"strconv"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_middleware "github.com/maximyunak/vetmessager/internal/core/transport/http/middleware"
	core_http_request "github.com/maximyunak/vetmessager/internal/core/transport/http/request"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
)

type CreateMessageRequest struct {
	ReplyToMessageID *int    `json:"reply_to_message_id"`
	Content          *string `json:"content" validate:"max=5000"`
}

func (h *MessagesHTTPHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke create message handler")

	userID, ok := core_http_middleware.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Unauthorized")
		return
	}
	chatID, err := strconv.Atoi(r.PathValue("chatId"))
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "Invalid chat ID")
	}

	var request CreateMessageRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"Failed to decode and validate CreateMessage request",
		)
		return
	}

	messageDomain := domainFromDTO(request, chatID, userID)

	messageDomain, err = h.messagesService.CreateMessage(ctx, messageDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create message")
		return
	}

	responseHandler.JSONResponse(messageDomain, http.StatusCreated)
}

func domainFromDTO(dto CreateMessageRequest, chatId int, senderId int) domain.Message {
	return domain.NewMessageUninitialized(
		chatId,
		senderId,
		dto.ReplyToMessageID,
		dto.Content,
	)
}

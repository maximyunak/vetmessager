package messages_transport_http

import (
	"net/http"
	"strconv"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_request "github.com/maximyunak/vetmessager/internal/core/transport/http/request"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
)

type CreateMessageRequest struct {
	ReplyToMessageID *int    `json:"reply_to_message_id"`
	Content          *string `json:"content" validate:"max=5000"`
}

// CreateMessage godoc
// @summary Create message
// @description Create a new message in the specified chat
// @tags messages
// @accept json
// @produce json
// @param chatId path int true "Chat ID"
// @param request body CreateMessageRequest true "Create message request"
// @success 201 {object} domain.Message "Message created successfully"
// @failure 400 {object} core_http_response.ErrorResponse "Invalid request or chat ID"
// @failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @failure 404 {object} core_http_response.ErrorResponse "Chat not found"
// @failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @router /chats/{chatId}/messages [post]
func (h *MessagesHTTPHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke create message handler")

	userID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Unauthorized")
		return
	}
	chatID, err := strconv.Atoi(r.PathValue("chatId"))
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "Invalid chat ID")
		return
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

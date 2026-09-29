package messages_transport_http

import (
	"net/http"
	"strconv"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
)

// DeleteMessage godoc
// @summary Delete message
// @description Delete an existing message from the specified chat
// @tags messages
// @produce json
// @param chatId path int true "Chat ID"
// @param messageId path int true "Message ID"
// @success 204 "Message deleted successfully"
// @failure 400 {object} core_http_response.ErrorResponse "Invalid chat ID or message ID"
// @failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @failure 404 {object} core_http_response.ErrorResponse "Message or chat not found"
// @failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @router /chats/{chatId}/messages/{messageId} [delete]
func (h *MessagesHTTPHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke delete message handler")

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
	messageID, err := strconv.Atoi(r.PathValue("messageId"))
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "Invalid message ID")
		return
	}

	if err := h.messagesService.DeleteMessage(ctx, userID, chatID, messageID); err != nil {
		responseHandler.ErrorResponse(err, "Error deleting message")
		return
	}

	responseHandler.JSONResponse(nil, http.StatusNoContent)

}

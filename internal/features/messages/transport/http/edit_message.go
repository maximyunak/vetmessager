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
	core_http_types "github.com/maximyunak/vetmessager/internal/core/transport/http/types"
)

type PatchMessageRequest struct {
	ReplyToMessageID core_http_types.Nullable[int]    `json:"reply_to_message_id"`
	Content          core_http_types.Nullable[string] `json:"content"`
}

// EditMessage GoDoc
// @summary Edit message
// @description Edit an existing message in the specified chat
// @tags messages
// @accept json
// @produce json
// @param chatId path int true "Chat ID"
// @param messageId path int true "Message ID"
// @param request body PatchMessageRequest true "Message update request"
// @success 200 {object} domain.Message "Message updated successfully"
// @failure 400 {object} core_http_response.ErrorResponse "Invalid request, chat ID, or message ID"
// @failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @failure 404 {object} core_http_response.ErrorResponse "Message or chat not found"
// @failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @router /chats/{chatId}/messages/{messageId} [patch]
func (h *MessagesHTTPHandler) EditMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke edit message handler")

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

	var request PatchMessageRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"Failed to decode and validate EditMessage request",
		)
		return
	}

	messagePatch := userPatchFromRequest(
		request,
		messageID,
		chatID,
		userID,
	)

	messageDomain, err := h.messagesService.EditMessage(ctx, messagePatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to update message")
		return
	}

	responseHandler.JSONResponse(messageDomain, http.StatusOK)
}
func userPatchFromRequest(
	request PatchMessageRequest,
	messageID int,
	chatID int,
	userID int,
) domain.MessagePatch {
	return domain.MessagePatch{
		ID:               messageID,
		ChatID:           chatID,
		SenderID:         userID,
		ReplyToMessageID: request.ReplyToMessageID.ToDomain(),
		Content:          request.Content.ToDomain(),
	}
}

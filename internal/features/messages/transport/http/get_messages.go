package messages_transport_http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
	core_http_utils "github.com/maximyunak/vetmessager/internal/core/transport/http/utils"
)

type GetMessagesResponse struct {
	Data   []MessageDTOResponse `json:"data"`
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Total  int                  `json:"total"`
}

// GetMessages godoc
// @summary Get chat messages
// @description Get messages from a chat with pagination
// @tags messages
// @produce json
// @param chatId path int true "Chat ID"
// @param limit query int false "Number of messages to return" default(20)
// @param offset query int false "Number of messages to skip" default(0)
// @success 200 {object} GetMessagesResponse "Messages retrieved successfully"
// @failure 400 {object} core_http_response.ErrorResponse "Invalid chat ID or pagination parameters"
// @failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @failure 404 {object} core_http_response.ErrorResponse "Chat not found"
// @failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @router /chats/{chatId}/messages [get]
func (h *MessagesHTTPHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke get users handler")

	userID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Unauthorized")
		return
	}

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit and offset parameter")

		return
	}

	chatID, err := strconv.Atoi(r.PathValue("chatId"))
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "Invalid chat ID")
		return
	}

	messageDomains, total, err := h.messagesService.GetMessages(ctx, userID, chatID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get messages")
		return
	}

	response := GetMessagesResponse{
		Data:   MessagesDTOFromDomain(messageDomains),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}

	responseHandler.JSONResponse(response, http.StatusOK)

}

func getLimitOffsetQueryParams(r *http.Request) (int, int, error) {
	limitParam, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return 0, 0, fmt.Errorf("get limit query param: %w", err)
	}

	offsetParam, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return 0, 0, fmt.Errorf("get offset query param: %w", err)
	}

	limit := 20
	offset := 0

	if limitParam != nil {
		limit = *limitParam
	}

	if offsetParam != nil {
		offset = *offsetParam
	}

	return limit, offset, nil
}

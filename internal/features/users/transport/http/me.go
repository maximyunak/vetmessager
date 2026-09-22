package users_transport_http

import (
	"net/http"

	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	log.Debug("invoke /me handler")

	userDomain, err := h.usersService.Me(ctx)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get me")
		return
	}

	responseHandler.JSONResponse(userDomain, http.StatusOK)

}

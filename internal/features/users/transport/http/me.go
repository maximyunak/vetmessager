package users_transport_http

import (
	"net/http"

	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
)

// Me godoc
// @summary Get current user
// @description Get the authenticated user's information
// @tags users
// @produce json
// @security BearerAuth
// @success 200 {object} UserDTOResponse "Success"
// @failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @failure 404 {object} core_http_response.ErrorResponse "User not found"
// @failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @router /me [get]
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

	response := UserDTOFromDomain(userDomain)

	responseHandler.JSONResponse(response, http.StatusOK)

}

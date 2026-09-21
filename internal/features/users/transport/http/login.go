package users_transport_http

import (
	"net/http"
	"time"

	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_request "github.com/maximyunak/vetmessager/internal/core/transport/http/request"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=5,max=100"`
	Password string `json:"password" validate:"required,min=5,max=100"`
}

type LoginResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *UsersHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke login user handler")

	var request LoginRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate LoginRequest")

		return
	}

	userDomain, err := h.usersService.Login(ctx, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to login")
		return
	}

	response := UserDTOFromDomain(userDomain)
	responseHandler.JSONResponse(response, http.StatusOK)
}

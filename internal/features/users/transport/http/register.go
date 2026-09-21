package users_transport_http

import (
	"net/http"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_request "github.com/maximyunak/vetmessager/internal/core/transport/http/request"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=5,max=255"`
	FirstName string `json:"first_name" validate:"required,min=5,max=255"`
	LastName  string `json:"last_name" validate:"required,min=5,max=255"`
	Email     string `json:"email" validate:"required,email,min=5,max=100"`
	Password  string `json:"password" validate:"required,min=5,max=100"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke create user handler")

	var request CreateUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate CreateUser request")

		return
	}

	userDomain := domainFromDTO(request)
	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")
		return
	}

	response := CreateUserResponse(UserDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.Username, dto.FirstName, dto.LastName, dto.Email, dto.Password)
}

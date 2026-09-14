package users_transport_http

import (
	"context"
	"net/http"

	"github.com/maximyunak/qzltgo/internal/core/domain"
	core_http_server "github.com/maximyunak/qzltgo/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	Login(
		ctx context.Context, email string, password string,
	) (domain.User, error)
	GetUsers(
		ctx context.Context, limit *int, offset *int,
	) ([]domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: h.Login,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
	}
}

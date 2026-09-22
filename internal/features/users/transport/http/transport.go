package users_transport_http

import (
	"context"
	"net/http"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_http_server "github.com/maximyunak/vetmessager/internal/core/transport/http/server"
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
	) (string, error)
	GetUsers(
		ctx context.Context, limit *int, offset *int,
	) ([]domain.User, error)
	Me(ctx context.Context) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService,
	}
}

func (h *UsersHTTPHandler) PublicRoutes() []core_http_server.Route {
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

func (h *UsersHTTPHandler) ProtectedRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/me",
			Handler: h.Me,
		},
	}
}

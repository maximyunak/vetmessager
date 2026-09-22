package users_service

import (
	"context"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type UsersService struct {
	UsersRepository UsersRepository
	PasswordHasher  PasswordHasher
	TokenManager    auth.TokenManager
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
}

func NewUsersService(usersRepository UsersRepository, passwordHasher PasswordHasher, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{
		UsersRepository: usersRepository,
		PasswordHasher:  passwordHasher,
		TokenManager:    tokenManager,
	}
}

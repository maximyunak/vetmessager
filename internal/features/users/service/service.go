package users_service

import (
	"context"

	"github.com/maximyunak/qzltgo/internal/core/domain"
)

type UsersService struct {
	UsersRepository UsersRepository
	PasswordHasher  PasswordHasher
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
}

func NewUsersService(usersRepository UsersRepository, passwordHasher PasswordHasher) *UsersService {
	return &UsersService{
		UsersRepository: usersRepository,
		PasswordHasher:  passwordHasher,
	}
}

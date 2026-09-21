package users_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

func (s *UsersService) Login(
	ctx context.Context,
	email string,
	password string,
) (domain.User, error) {
	user, err := s.UsersRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	isCorrectPassword, err := s.PasswordHasher.Compare(user.Password, password)

	if err != nil {
		return domain.User{}, fmt.Errorf("compare password: %w", err)
	}

	if !isCorrectPassword {
		return domain.User{}, fmt.Errorf("invalid password")
	}

	return user, nil
}

package users_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/qzltgo/internal/core/domain"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, err
	}

	hash, err := s.PasswordHasher.Hash(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hashing password: %w", err)
	}

	user.Password = hash

	user, err = s.UsersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

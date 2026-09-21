package users_service

import (
	"context"
	"fmt"
)

func (s *UsersService) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	user, err := s.UsersRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	isCorrectPassword, err := s.PasswordHasher.Compare(user.Password, password)

	if err != nil {
		return "", fmt.Errorf("compare password: %w", err)
	}

	if !isCorrectPassword {
		return "", fmt.Errorf("invalid password")
	}

	accessToken, err := s.TokenManager.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return accessToken, nil
}

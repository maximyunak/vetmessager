package users_service

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_http_middleware "github.com/maximyunak/vetmessager/internal/core/transport/http/middleware"
)

func (s *UsersService) Me(ctx context.Context) (domain.User, error) {
	userId, ok := core_http_middleware.UserIDFromContext(ctx)

	if !ok {
		return domain.User{}, fmt.Errorf("token not found")
	}
	user, err := s.UsersRepository.GetUserById(ctx, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %v: %w", err, core_errors.ErrNotFound)
	}

	return user, nil
}

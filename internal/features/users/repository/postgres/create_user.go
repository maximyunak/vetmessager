package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/maximyunak/qzltgo/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO users (name, email, password) values ($1, $2, $3) RETURNING id, name, email, password,created_at,updated_at 
	`
	row := r.pool.QueryRow(ctx, query, user.Name, user.Email, user.Password)

	var userModel UserModel
	err := row.Scan(&userModel.ID, &userModel.Name, &userModel.Email, &userModel.Password, &userModel.CreatedAt, &userModel.UpdatedAt)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(userModel.ID, userModel.Name, userModel.Email, userModel.Password)

	return userDomain, nil
}

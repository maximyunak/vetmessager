package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO users (username, first_name, last_name, email, password, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id, username,first_name,last_name, email, password,created_at,updated_at 
	`
	row := r.pool.QueryRow(ctx, query, user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	var userModel UserModel
	err := row.Scan(&userModel.ID, &userModel.UserName, &userModel.FirstName, &userModel.LastName, &userModel.Email, &userModel.Password, &userModel.CreatedAt, &userModel.UpdatedAt)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := userDomainFromModel(userModel)

	return userDomain, nil
}

package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, username, first_name, last_name, email, password, created_at, updated_at FROM users WHERE email = $1
	`
	row := r.pool.QueryRow(ctx, query, email)

	var userModel UserModel
	err := row.Scan(&userModel.ID, &userModel.UserName, &userModel.FirstName, &userModel.LastName, &userModel.Email, &userModel.Password, &userModel.CreatedAt, &userModel.UpdatedAt)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(userModel.ID, userModel.UserName, userModel.FirstName, userModel.LastName, userModel.Email, userModel.Password, userModel.CreatedAt, userModel.UpdatedAt)

	return userDomain, nil
}

func (r *UsersRepository) GetUserById(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, username, first_name, last_name, email, password, created_at, updated_at FROM users WHERE id = $1
	`
	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel
	err := row.Scan(&userModel.ID, &userModel.UserName, &userModel.FirstName, &userModel.LastName, &userModel.Email, &userModel.Password, &userModel.CreatedAt, &userModel.UpdatedAt)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(userModel.ID, userModel.UserName, userModel.FirstName, userModel.LastName, userModel.Email, userModel.Password, userModel.CreatedAt, userModel.UpdatedAt)

	return userDomain, nil
}

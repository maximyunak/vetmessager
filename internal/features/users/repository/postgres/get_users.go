package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/maximyunak/qzltgo/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, name, email, password, created_at, updated_at 
	FROM users 
	ORDER BY id ASC
	LIMIT $1 
	OFFSET $2 
	`
	rows, err := r.pool.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("postgres query: %w", err)
	}
	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel

		err := rows.Scan(&userModel.ID, &userModel.Name, &userModel.Email, &userModel.Password, &userModel.CreatedAt, &userModel.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("postgres scan: %w", err)
		}

		userModels = append(userModels, userModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres rows: %w", err)
	}

	userDomains := userDomainsFromModels(userModels)
	return userDomains, nil
}

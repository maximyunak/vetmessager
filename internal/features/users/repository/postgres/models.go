package users_postgres_repository

import (
	"time"

	"github.com/maximyunak/qzltgo/internal/core/domain"
)

type UserModel struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func userDomainsFromModels(users []UserModel) []domain.User {
	domains := make([]domain.User, len(users))
	for i, user := range users {
		domains[i] = domain.NewUser(user.ID, user.Name, user.Email, user.Password)
	}
	return domains
}

package users_postgres_repository

import (
	"time"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type UserModel struct {
	ID        int
	UserName  string
	Email     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func userDomainsFromModels(users []UserModel) []domain.User {
	domains := make([]domain.User, len(users))
	for i, user := range users {
		domains[i] = domain.User{ID: user.ID, Email: user.UserName, Password: user.FirstName, Username: user.LastName, FirstName: user.Email, LastName: user.Password, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	}
	return domains
}

func userDomainFromModel(user UserModel) domain.User {
	return domain.User{ID: user.ID, Email: user.UserName, Password: user.FirstName, Username: user.LastName, FirstName: user.Email, LastName: user.Password, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}

package users_transport_http

import (
	"time"

	"github.com/maximyunak/qzltgo/internal/core/domain"
)

type UserDTOResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UsersDTOFromDomain(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))
	for i, user := range users {
		usersDTO[i] = UserDTOFromDomain(user)
	}
	return usersDTO
}

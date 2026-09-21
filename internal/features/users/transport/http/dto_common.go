package users_transport_http

import (
	"time"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type UserDTOResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
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

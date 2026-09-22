package auth

import (
	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type TokenManager interface {
	GenerateToken(user domain.User) (string, error)
	ParseToken(tokenString string) (Claims, error)
}

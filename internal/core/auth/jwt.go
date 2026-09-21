package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maximyunak/vetmessager/internal/core/domain"
)

type JWTManager struct {
	secret []byte
	ttl    time.Duration
}

func NewJWTManager(config Config) *JWTManager {
	return &JWTManager{secret: config.Secret, ttl: config.AccessTTL}
}

func (m *JWTManager) GenerateToken(user domain.User) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.secret)
}

func (m *JWTManager) ParseToken(tokenString string) (Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrTokenSignatureInvalid
			}

			return m.secret, nil
		},
	)

	if err != nil {
		return Claims{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return Claims{}, jwt.ErrTokenInvalidClaims
	}

	return *claims, nil
}

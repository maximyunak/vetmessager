package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/maximyunak/qzltgo/internal/core/errors"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(id int, name string, email string, password string) User {
	return User{
		ID:       id,
		Email:    email,
		Name:     name,
		Password: password,
	}
}

func NewUserUninitialized(name string, email string, password string) User {
	return NewUser(UninitializedID, name, email, password)
}

func (u User) Validate() error {
	nameLength := len([]rune(u.Name))
	if nameLength < 3 || nameLength > 100 {
		return fmt.Errorf("invalid `name` length: %d", nameLength, core_errors.ErrInvalidArgument)
	}

	passwordLength := len([]rune(u.Password))
	if passwordLength < 3 || passwordLength > 100 {
		return fmt.Errorf("invalid `name` length: %d", passwordLength, core_errors.ErrInvalidArgument)
	}

	emailLength := len([]rune(u.Email))
	if emailLength < 3 || emailLength > 100 {
		return fmt.Errorf("invalid `name` length: %d", emailLength, core_errors.ErrInvalidArgument)
	}

	emailRe := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !emailRe.MatchString(u.Email) {
		return fmt.Errorf("invalid `email` format: %s", u.Email)
	}

	return nil
}

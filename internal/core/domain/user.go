package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Username  string    `json:"username" db:"username"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(id int, username string, firstName string, lastName string, email string, password string, createdAt time.Time, updatedAt time.Time) User {
	return User{
		ID:        id,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewUserUninitialized(username string, firstName string, lastName string, email string, password string) User {
	return NewUser(UninitializedID, username, firstName, lastName, email, password, time.Now(), time.Now())
}

func (u User) Validate() error {
	usernameLength := len([]rune(u.Username))
	if usernameLength < 3 || usernameLength > 100 {
		return fmt.Errorf("invalid `username` length: %d", usernameLength, core_errors.ErrInvalidArgument)
	}

	firstNameLength := len([]rune(u.FirstName))
	if firstNameLength < 3 || firstNameLength > 100 {
		return fmt.Errorf("invalid `first_name` length: %d", firstNameLength, core_errors.ErrInvalidArgument)
	}

	lastNameLength := len([]rune(u.LastName))
	if lastNameLength < 3 || lastNameLength > 100 {
		return fmt.Errorf("invalid `last_name` length: %d", lastNameLength, core_errors.ErrInvalidArgument)
	}

	passwordLength := len([]rune(u.Password))
	if passwordLength < 3 || passwordLength > 100 {
		return fmt.Errorf("invalid `password` length: %d", passwordLength, core_errors.ErrInvalidArgument)
	}

	emailLength := len([]rune(u.Email))
	if emailLength < 3 || emailLength > 100 {
		return fmt.Errorf("invalid `email` length: %d", emailLength, core_errors.ErrInvalidArgument)
	}

	emailRe := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !emailRe.MatchString(u.Email) {
		return fmt.Errorf("invalid `email` format: %s", u.Email)
	}

	return nil
}

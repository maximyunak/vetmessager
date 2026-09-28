package domain

import (
	"fmt"
	"time"

	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

type Chat struct {
	ID        int       `json:"id" db:"id"`
	Name      *string   `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	CreatorID *int      `json:"creator_id" db:"creator_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewChatUninitialized(name *string, chatType string, creatorID *int) Chat {
	return Chat{
		ID:        UninitializedID,
		Name:      name,
		Type:      chatType,
		CreatorID: creatorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c Chat) Validate() error {
	if c.Type != "personal" && c.Type != "group" {
		return fmt.Errorf(
			"invalid `type`: %s: %w",
			c.Type,
			core_errors.ErrInvalidArgument,
		)
	}

	if c.CreatorID != nil && *c.CreatorID <= 0 {
		return fmt.Errorf(
			"invalid `creator_id`: %d: %w",
			*c.CreatorID,
			core_errors.ErrInvalidArgument,
		)
	}

	if c.Name != nil && len([]rune(*c.Name)) > 255 {
		return fmt.Errorf(
			"invalid `name` length: %d: %w",
			len([]rune(*c.Name)),
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

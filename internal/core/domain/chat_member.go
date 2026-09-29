package domain

import "time"

type ChatMember struct {
	ChatID    int       `json:"chat_id" db:"chat_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	IsMuted   bool      `json:"is_muted" db:"is_muted"`
}

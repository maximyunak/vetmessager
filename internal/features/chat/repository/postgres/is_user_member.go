package chat_postgres_repository

import (
	"context"
	"fmt"
	"time"

	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

func (r *ChatRepository) IsUserMember(ctx context.Context, chatID int, userID int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
	SELECT EXISTS (
		SELECT 1
    	FROM chat_members
    	WHERE chat_id = $1 AND user_id = $2
	)
	`

	var exists bool

	err := r.pool.QueryRow(ctx, query, chatID, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check chat membership: %v: %w", err, core_errors.ErrNotFound)
	}

	return exists, nil
}

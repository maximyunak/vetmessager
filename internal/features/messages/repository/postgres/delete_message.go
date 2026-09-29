package messages_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

func (r *MessageRepository) DeleteMessage(
	ctx context.Context,
	messageID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE messages
		SET deleted_at = CURRENT_TIMESTAMP,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`

	result, err := r.pool.Exec(
		ctx,
		query,
		messageID,
	)

	if err != nil {
		return fmt.Errorf("delete message: %w", err)
	}

	if result.RowsAffected() == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}

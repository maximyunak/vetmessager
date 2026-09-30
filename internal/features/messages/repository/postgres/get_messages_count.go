package messages_postgres_repository

import (
	"context"
	"fmt"
)

func (r *MessageRepository) GetMessagesCount(
	ctx context.Context,
	chatID int,
) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT COUNT(*) FROM messages WHERE chat_id = $1 AND deleted_at IS NULL;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		chatID,
	)

	var count int

	err := row.Scan(
		&count,
	)
	if err != nil {
		return 0, fmt.Errorf("get messages count: %w", err)
	}

	return count, nil
}

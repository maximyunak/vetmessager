package chat_postgres_repository

import (
	"context"
	"fmt"
)

func (r *ChatRepository) GetUserChatsIDs(ctx context.Context, userID int) ([]int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT chat_id FROM chat_members WHERE user_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, userID)

	if err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}
	defer rows.Close()
	var chatIDs []int

	for rows.Next() {
		var chatID int

		err := rows.Scan(&chatID)
		if err != nil {
			return nil, fmt.Errorf("postgres scan: %w", err)
		}

		chatIDs = append(chatIDs, chatID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return chatIDs, nil
}

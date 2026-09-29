package messages_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

func (r *MessageRepository) EditMessage(
	ctx context.Context,
	message domain.Message,
) (domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
        UPDATE messages
        SET reply_to_message_id = $1,
            content = $2,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $3
        RETURNING id, chat_id, sender_id, reply_to_message_id,
                  content, created_at, updated_at, deleted_at
    `

	err := r.pool.QueryRow(
		ctx,
		query,
		message.ReplyToMessageID,
		message.Content,
		message.ID,
	).Scan(
		&message.ID,
		&message.ChatID,
		&message.SenderID,
		&message.ReplyToMessageID,
		&message.Content,
		&message.CreatedAt,
		&message.UpdatedAt,
		&message.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Message{}, core_errors.ErrNotFound
		}

		return domain.Message{}, fmt.Errorf("edit message: %w", err)
	}

	return message, nil
}

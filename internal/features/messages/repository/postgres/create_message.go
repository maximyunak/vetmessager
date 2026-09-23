package messages_postgres_repository

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
)

func (r *MessageRepository) CreateMessage(
	ctx context.Context,
	messageDomain domain.Message,
) (domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO messages (
			chat_id,
			sender_id,
			reply_to_message_id,
			content,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			chat_id,
			sender_id,
			reply_to_message_id,
			content,
			created_at,
			updated_at,
			deleted_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		messageDomain.ChatID,
		messageDomain.SenderID,
		messageDomain.ReplyToMessageID,
		messageDomain.Content,
		messageDomain.CreatedAt,
		messageDomain.UpdatedAt,
	)

	var messageModel MessageModel

	err := row.Scan(
		&messageModel.ID,
		&messageModel.ChatID,
		&messageModel.SenderID,
		&messageModel.ReplyToMessageID,
		&messageModel.Content,
		&messageModel.CreatedAt,
		&messageModel.UpdatedAt,
		&messageModel.DeletedAt,
	)
	if err != nil {
		return domain.Message{}, fmt.Errorf("scan error: %w", err)
	}

	messageDomain = messageDomainFromModel(messageModel)

	return messageDomain, nil
}

package messages_postgres_repository

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

func (r *MessageRepository) GetMessages(
	ctx context.Context,
	chatID int,
	limit int,
	offset int,
) ([]domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT
		id,
		chat_id,
		sender_id,
		reply_to_message_id,
		content,
		created_at,
		updated_at,
		deleted_at
	FROM messages
	WHERE chat_id = $1
	  AND deleted_at IS NULL
	ORDER BY created_at DESC
	LIMIT $2
	OFFSET $3;
`

	rows, err := r.pool.Query(
		ctx,
		query,
		chatID,
		limit,
		offset,
	)
	defer rows.Close()

	var messagesModel []MessageModel

	for rows.Next() {
		var messageModel MessageModel

		err := rows.Scan(
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
			return nil, fmt.Errorf("scan message: %w", err)
		}

		messagesModel = append(messagesModel, messageModel)
	}

	if err != nil {
		return nil, fmt.Errorf("insert message: %v: %w", err, core_errors.ErrNotFound)
	}

	messagesDomain := messageDomainsFromModels(messagesModel)

	return messagesDomain, nil
}

package messages_postgres_repository

import (
	"context"
	"fmt"

	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

func (r *MessageRepository) GetMessage(
	ctx context.Context,
	messageID int,
) (domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT * FROM messages WHERE id = $1;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		messageID,
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
		return domain.Message{}, fmt.Errorf("insert message: %v: %w", err, core_errors.ErrNotFound)
	}

	messageDomain := messageDomainFromModel(messageModel)

	return messageDomain, nil
}

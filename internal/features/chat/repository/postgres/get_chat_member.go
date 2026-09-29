package chat_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/maximyunak/vetmessager/internal/core/domain"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
)

func (r *ChatRepository) GetChatMember(
	ctx context.Context,
	chatID int,
	userID int,
) (domain.ChatMember, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		SELECT chat_id, user_id, role, created_at, updated_at
		FROM chat_members
		WHERE chat_id = $1 AND user_id = $2
	`

	var model ChatMemberModel

	err := r.pool.QueryRow(
		ctx,
		query,
		chatID,
		userID,
	).Scan(
		&model.ChatID,
		&model.UserID,
		&model.Role,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if err != nil {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return domain.ChatMember{}, core_errors.ErrNotFound
			}

			return domain.ChatMember{}, fmt.Errorf("get chat member: %w", err)
		}
	}

	return chatMemberDomainFromModel(model), nil
}

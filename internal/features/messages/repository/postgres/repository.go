package messages_postgres_repository

import core_postgres_pool "github.com/maximyunak/vetmessager/internal/core/repository/postgres/pull"

type MessageRepository struct {
	pool core_postgres_pool.Pool
}

func NewMessagesRepository(pool core_postgres_pool.Pool) *MessageRepository {
	return &MessageRepository{pool: pool}
}

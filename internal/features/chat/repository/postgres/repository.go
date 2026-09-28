package chat_postgres_repository

import core_postgres_pool "github.com/maximyunak/vetmessager/internal/core/repository/postgres/pull"

type ChatRepository struct {
	pool core_postgres_pool.Pool
}

func NewChatRepository(pool core_postgres_pool.Pool) *ChatRepository {
	return &ChatRepository{
		pool: pool,
	}
}

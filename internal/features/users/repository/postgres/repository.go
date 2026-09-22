package users_postgres_repository

import core_postgres_pool "github.com/maximyunak/vetmessager/internal/core/repository/postgres/pull"

type UsersRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(pool core_postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}

package postgresql

import "github.com/jackc/pgx/v4/pgxpool"

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{db: db}
}

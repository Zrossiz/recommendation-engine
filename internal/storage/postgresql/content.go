package postgresql

import "github.com/jackc/pgx/v4/pgxpool"

type ContentStore struct {
	db *pgxpool.Pool
}

func NewContentStore(db *pgxpool.Pool) *ContentStore {
	return &ContentStore{db: db}
}

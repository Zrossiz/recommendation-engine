package postgresql

import "github.com/jackc/pgx/v4/pgxpool"

type CategoryStore struct {
	db *pgxpool.Pool
}

func NewCategoryStore(db *pgxpool.Pool) *CategoryStore {
	return &CategoryStore{db: db}
}

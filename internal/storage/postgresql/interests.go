package postgresql

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type InterestsStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewInterestsStore(db *pgxpool.Pool, log *zap.Logger) *InterestsStore {
	return &InterestsStore{
		db:  db,
		log: log,
	}
}

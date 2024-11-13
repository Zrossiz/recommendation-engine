package postgresql

import (
	"context"
	"engine/internal/dto"

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

func (i *InterestsStore) Create(interestDTO dto.Interest) (bool, error) {
	sql := `INSERT INTO interests (user_id, category_id) VALUES ($1, $2)`
	_, err := i.db.Exec(
		context.Background(),
		sql,
		interestDTO.UserId,
		interestDTO.CategoryId,
	)
	if err != nil {
		i.log.Error("create interest error", zap.Error(err))
		return false, err
	}

	return true, nil
}

package postgresql

import (
	"engine/internal/dto"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type UserInteractionsStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewUserInteractionsStore(db *pgxpool.Pool, log *zap.Logger) *UserInteractionsStore {
	return &UserInteractionsStore{
		db:  db,
		log: log,
	}
}

func Create(dto dto.CreateInteraction) (bool, error) {
	return true, nil
}

func GetInteractionsByUser(userID int64) (dto.GetUserInteractions, error) {
	var resp dto.GetUserInteractions
	return resp, nil
}

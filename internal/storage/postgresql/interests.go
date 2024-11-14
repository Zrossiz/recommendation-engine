package postgresql

import (
	"context"
	"engine/internal/dto"
	"engine/internal/model"

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

func (i *InterestsStore) GetUserInterests(userID int64) ([]model.Interest, error) {
	sql := `SELECT id, user_id, category_id FROM interests WHERE user_id = $1`

	rows, err := i.db.Query(context.Background(), sql, userID)
	if err != nil {
		i.log.Error("error query GetUserInterests", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var interests []model.Interest
	for rows.Next() {
		var interest model.Interest
		err := rows.Scan(
			&interest.ID,
			&interest.UserID,
			&interest.CategoryID,
		)
		if err != nil {
			i.log.Error("GetUserInterests scan", zap.Error(err))
			return nil, err
		}

		interests = append(interests, interest)
	}

	if err = rows.Err(); err != nil {
		i.log.Error("rows error GetUserInterests", zap.Error(err))
		return nil, err
	}

	return interests, nil
}

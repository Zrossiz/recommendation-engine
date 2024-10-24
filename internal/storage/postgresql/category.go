package postgresql

import (
	"context"
	"engine/internal/apperrors"
	"engine/internal/dto"
	"engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type CategoryStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewCategoryStore(db *pgxpool.Pool, log *zap.Logger) *CategoryStore {
	return &CategoryStore{
		db:  db,
		log: log,
	}
}

func (c *CategoryStore) Create(categoryDTO dto.Category) (bool, error) {
	sql := `INSERT INTO categories (name) VALUES ($1)`
	_, err := c.db.Exec(context.Background(), sql, categoryDTO.Name)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *CategoryStore) GetAll() ([]model.Category, error) {
	sql := `SELECT id, name FROM categories`
	rows, err := c.db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryStore) Delete(id int64) (bool, error) {
	sql := `DELETE FROM categories WHERE id = $1`
	cmdTag, err := c.db.Exec(context.Background(), sql, id)
	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, apperrors.ErrNotFound
	}

	return true, nil
}

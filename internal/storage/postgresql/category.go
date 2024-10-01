package postgresql

import (
	"engine/internal/dto"
	models "engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type CategoryStore struct {
	db *pgxpool.Pool
}

func NewCategoryStore(db *pgxpool.Pool) *CategoryStore {
	return &CategoryStore{db: db}
}

func (c *CategoryStore) Create(categoryDTO dto.Category) (bool, error) {
	return true, nil
}

func (c *CategoryStore) GetAll() ([]models.Category, error) {
	var categories []models.Category
	return categories, nil
}

func (c *CategoryStore) Delete(id int64) (bool, error) {
	return true, nil
}

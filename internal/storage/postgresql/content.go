package postgresql

import (
	"engine/internal/dto"
	"engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ContentStore struct {
	db *pgxpool.Pool
}

func NewContentStore(db *pgxpool.Pool) *ContentStore {
	return &ContentStore{db: db}
}

func (co *ContentStore) Create(contentDTO dto.Content) (bool, error) {
	return true, nil
}

func (co *ContentStore) Delete(id int64) (bool, error) {
	return true, nil
}

func (co *ContentStore) Update(contentDTO dto.Content) (bool, error) {
	return true, nil
}

func (co *ContentStore) GetAllByCategory(categoryId int64, page int64) ([]model.Content, error) {
	var content []model.Content

	return content, nil
}

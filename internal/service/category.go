package service

import (
	"engine/internal/dto"
	"engine/internal/model"

	"go.uber.org/zap"
)

type CategoryStore interface {
	Create(categoryDTO dto.Category) (bool, error)
	Delete(id int64) (bool, error)
	GetAll() ([]model.Category, error)
}

type CategoryService struct {
	db  CategoryStore
	log *zap.Logger
}

func NewCategoryService(db CategoryStore, log *zap.Logger) *CategoryService {
	return &CategoryService{
		db:  db,
		log: log,
	}
}

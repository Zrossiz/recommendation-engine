package service

import (
	"engine/internal/dto"
	"engine/internal/model"

	"go.uber.org/zap"
)

type ContentStore interface {
	Create(contentDTO dto.Content) (bool, error)
	Delete(id int64) (bool, error)
	GetAllByCategory(categoryId int64, page int64) ([]model.Content, error)
	Update(contentID int64, contentDTO dto.Content) (bool, error)
}

type ContentService struct {
	db  ContentStore
	log *zap.Logger
}

func NewContentService(db ContentStore, log *zap.Logger) *ContentService {
	return &ContentService{
		db:  db,
		log: log,
	}
}

func (c *ContentService) Create(contentDTO dto.Content) (bool, error) {
	_, err := c.db.Create(contentDTO)
	if err != nil {
		return false, err
	}

	return true, nil
}

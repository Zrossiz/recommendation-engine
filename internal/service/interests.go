package service

import (
	"engine/internal/dto"
	"engine/internal/model"

	"go.uber.org/zap"
)

type InterestsService struct {
	db  InterestsStore
	log *zap.Logger
}

type InterestsStore interface {
	Create(interestDTO dto.Interest) (bool, error)
	GetUserInterests(userID int64) ([]model.Interest, error)
}

func NewInterestsService(db InterestsStore, log *zap.Logger) *InterestsService {
	return &InterestsService{
		db:  db,
		log: log,
	}
}

func (i *InterestsService) Create(interestDTO dto.Interest) (bool, error) {
	return i.db.Create(interestDTO)
}

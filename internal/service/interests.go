package service

import (
	"engine/internal/dto"

	"go.uber.org/zap"
)

type InterestsService struct {
	db  InterestsStore
	log *zap.Logger
}

type InterestsStore interface {
	Create(interestDTO dto.Interest) (bool, error)
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

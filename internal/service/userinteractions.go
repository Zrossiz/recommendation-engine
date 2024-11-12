package service

import (
	"engine/internal/dto"

	"go.uber.org/zap"
)

type UserInteractionsStore interface {
	Create(dto dto.CreateInteraction) (bool, error)
	GetInteractionsByUser(userID int64) ([]dto.GetUserInteractions, error)
}

type UserInteractionsService struct {
	db  UserInteractionsStore
	log *zap.Logger
}

func NewUserInteractionsService(db UserInteractionsStore, log *zap.Logger) *UserInteractionsService {
	return &UserInteractionsService{
		db:  db,
		log: log,
	}
}

func (i *UserInteractionsService) Create(interactionDTO dto.CreateInteraction) (bool, error) {
	_, err := i.db.Create(interactionDTO)
	if err != nil {
		return false, err
	}

	return true, nil
}

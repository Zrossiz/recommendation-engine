package service

import "go.uber.org/zap"

type InterestsService struct {
	db  InterestsStore
	log *zap.Logger
}

type InterestsStore interface {
}

func NewInterestsService(db InterestsStore, log *zap.Logger) *InterestsService {
	return &InterestsService{
		db:  db,
		log: log,
	}
}

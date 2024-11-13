package service

import (
	"engine/internal/config"

	"go.uber.org/zap"
)

type Service struct {
	Category         *CategoryService
	User             *UserService
	Content          *ContentService
	UserInteractions *UserInteractionsService
	Interests        *InterestsService
}

type Storage struct {
	UserStore             UserStore
	CategoryStore         CategoryStore
	ContentStore          ContentStore
	UserInteractionsStore UserInteractionsStore
	InterestsStore        InterestsStore
}

func New(db Storage, log *zap.Logger, cfg *config.Config) *Service {
	return &Service{
		Category:         NewCategoryService(db.CategoryStore, log),
		User:             NewUserService(db.UserStore, log, cfg),
		Content:          NewContentService(db.ContentStore, db.UserInteractionsStore, log),
		UserInteractions: NewUserInteractionsService(db.UserInteractionsStore, log),
		Interests:        NewInterestsService(db.InterestsStore, log),
	}
}

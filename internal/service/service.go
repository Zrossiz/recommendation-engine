package service

import (
	"go.uber.org/zap"
)

type Service struct {
	Category         *CategoryService
	User             *UserService
	Content          *ContentService
	UserInteractions *UserInteractionsService
}

type Storage struct {
	UserStore             UserStore
	CategoryStore         CategoryStore
	ContentStore          ContentStore
	UserInteractionsStore UserInteractionsStore
}

func New(db Storage, log *zap.Logger) *Service {
	return &Service{
		Category:         NewCategoryService(db.CategoryStore, log),
		User:             NewUserService(db.UserStore, log),
		Content:          NewContentService(db.ContentStore, log),
		UserInteractions: NewUserInteractionsService(db.UserInteractionsStore, log),
	}
}

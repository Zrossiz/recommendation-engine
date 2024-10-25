package service

import (
	"engine/internal/dto"
	"engine/internal/model"

	"go.uber.org/zap"
)

type UserStore interface {
	Create(userDTO dto.User) (bool, error)
	Delete(id int64) (bool, error)
	GetUserById(id int64) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
	Update(id int64, userDTO dto.User) (bool, error)
}

type UserService struct {
	db  UserStore
	log *zap.Logger
}

func NewUserService(db UserStore, log *zap.Logger) *UserService {
	return &UserService{
		db:  db,
		log: log,
	}
}

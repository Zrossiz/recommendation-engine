package postgresql

import (
	"engine/internal/dto"
	models "engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{db: db}
}

func (u *UserStore) Create(userDTO dto.User) (bool, error) {
	return true, nil
}

func (u *UserStore) Delete(id int64) (bool, error) {
	return true, nil
}

func (u *UserStore) Update(id int64, userDTO dto.User) (*models.User, error) {
	return nil, nil
}

func (u *UserStore) GetUserById(id int64) (*models.User, error) {
	return nil, nil
}

func (u *UserStore) GetUserByName(name string) (*models.User, error) {
	return nil, nil
}

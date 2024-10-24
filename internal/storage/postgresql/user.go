package postgresql

import (
	"context"
	"engine/internal/apperrors"
	"engine/internal/dto"
	"engine/internal/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type UserStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewUserStore(db *pgxpool.Pool, log *zap.Logger) *UserStore {
	return &UserStore{
		db:  db,
		log: log,
	}
}

func (u *UserStore) Create(userDTO dto.User) (bool, error) {
	sql := `INSERT INTO users (name, password) VALUES ($1, $2)`
	_, err := u.db.Exec(context.Background(), sql, userDTO.Username, userDTO.Password)
	if err != nil {
		u.log.Error("insert user", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (u *UserStore) Delete(id int64) (bool, error) {
	sql := `DELETE FROM users WHERE id = $1`
	cmdTag, err := u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, apperrors.ErrNotFound
	}

	return true, nil
}

func (u *UserStore) Update(id int64, userDTO dto.User) (bool, error) {
	sql := `UPDATE users SET "name" = $1, password = $2 WHERE id = $3`
	cmdTag, err := u.db.Exec(context.Background(), sql, userDTO.Username, userDTO.Password, id)
	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}

func (u *UserStore) GetUserById(id int64) (*model.User, error) {
	sql := `SELECT id, name FROM users WHERE id = $1`

	row := u.db.QueryRow(context.Background(), sql, id)
	var usr model.User
	err := row.Scan(&usr.ID, &usr.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		u.log.Error("get user by id", zap.Error(err))
		return nil, err
	}

	return &usr, nil
}

func (u *UserStore) GetUserByName(name string) (*model.User, error) {
	sql := `SELECT id, name FROM users WHERE "name" = $1`

	row := u.db.QueryRow(context.Background(), sql, name)
	var usr model.User
	err := row.Scan(&usr.ID, &usr.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		u.log.Error("get user by id", zap.Error(err))
		return nil, err
	}

	return &usr, nil
}

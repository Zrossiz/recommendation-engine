package postgresql

import (
	"context"
	"engine/internal/apperrors"
	"engine/internal/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type TokenStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewTokenStore(db *pgxpool.Pool, log *zap.Logger) *TokenStore {
	return &TokenStore{
		db:  db,
		log: log,
	}
}

func (t *TokenStore) Create(userID int64, token string) (bool, error) {
	sql := `INSERT INTO tokens (user_id, token) VALUES ($1, $2)`
	_, err := t.db.Exec(context.Background(), sql, userID, token)
	if err != nil {
		t.log.Error("insert token", zap.Error(err))
		return false, err
	}
	return true, nil
}

func (t *TokenStore) DeleteByToken(token string) (bool, error) {
	sql := `DELETE FROM tokens WHERE token = $1`
	cmdTag, err := t.db.Exec(context.Background(), sql, token)
	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, apperrors.ErrNotFound
	}

	return true, nil
}

func (t *TokenStore) DeleteTokensByUser(userID int64) (bool, error) {
	sql := `DELETE FROM tokens WHERE user_id = $1`
	cmdTag, err := t.db.Exec(context.Background(), sql, userID)
	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, apperrors.ErrNotFound
	}

	return true, nil
}

func (t *TokenStore) GetTokenByToken(token string) (*model.RefreshToken, error) {
	sql := `SELECT id, user_id, token FROM tokens WHERE token = $1`
	row := t.db.QueryRow(context.Background(), sql, token)

	var tkn model.RefreshToken
	err := row.Scan(&tkn.ID, &tkn.UserID, &tkn.Token)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		t.log.Error("get token by token", zap.Error(err))
		return nil, err
	}
	return &tkn, nil
}

func (t *TokenStore) GetTokenByUser(userID int64) (*model.RefreshToken, error) {
	sql := `SELECT id, user_id, token FROM tokens WHERE user_id = $1`
	row := t.db.QueryRow(context.Background(), sql, userID)

	var tkn model.RefreshToken
	err := row.Scan(&tkn.ID, &tkn.UserID, &tkn.Token)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		t.log.Error("get token by token", zap.Error(err))
		return nil, err
	}
	return &tkn, nil
}

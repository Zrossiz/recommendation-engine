package postgresql

import (
	models "engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TokenStore struct {
	db *pgxpool.Pool
}

func NewTokenStore(db *pgxpool.Pool) *TokenStore {
	return &TokenStore{db: db}
}

func (t *TokenStore) Create(userID int64, token string) (bool, error) {
	return true, nil
}

func (t *TokenStore) DeleteByToken(token string) (bool, error) {
	return true, nil
}

func (t *TokenStore) DeleteTokensByUsers(userID int64) (bool, error) {
	return true, nil
}

func (t *TokenStore) GetTokenByToken(token string) (*models.RefreshToken, error) {
	return nil, nil
}

func (t *TokenStore) GetTokenByUser(userID int64) (*models.RefreshToken, error) {
	return nil, nil
}

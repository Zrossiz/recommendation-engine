package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type DBStorage struct {
	ContentStore          *ContentStore
	CategoryStore         *CategoryStore
	UserStore             *UserStore
	TokenStore            *TokenStore
	CommentStore          *CommentStore
	UserInteractionsStore *UserInteractionsStore
}

func Connect(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("connect fail")
	}

	return db, nil
}

func New(dbConn *pgxpool.Pool, log *zap.Logger) *DBStorage {

	return &DBStorage{
		ContentStore:          NewContentStore(dbConn, log),
		UserStore:             NewUserStore(dbConn, log),
		CategoryStore:         NewCategoryStore(dbConn, log),
		TokenStore:            NewTokenStore(dbConn, log),
		CommentStore:          NewCommentStore(dbConn, log),
		UserInteractionsStore: NewUserInteractionsStore(dbConn, log),
	}
}

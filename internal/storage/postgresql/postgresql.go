package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type DBStorage struct {
	ContentStore          *ContentStore
	CategoryStore         *CategoryStore
	UserStore             *UserStore
	UserInteractionsStore *UserInteractionsStore
}

func Connect(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func New(dbConn *pgxpool.Pool, log *zap.Logger) *DBStorage {

	return &DBStorage{
		ContentStore:          NewContentStore(dbConn, log),
		UserStore:             NewUserStore(dbConn, log),
		CategoryStore:         NewCategoryStore(dbConn, log),
		UserInteractionsStore: NewUserInteractionsStore(dbConn, log),
	}
}

package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBStorage struct {
	ContentStore  *ContentStore
	CategoryStore *CategoryStore
	UserStore     *UserStore
	TokenStore    *TokenStore
	CommentStore  *CommentStore
}

func Connect(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("connect fail")
	}

	return db, nil
}

func New(dbConn *pgxpool.Pool) *DBStorage {

	return &DBStorage{
		ContentStore:  NewContentStore(dbConn),
		UserStore:     NewUserStore(dbConn),
		CategoryStore: NewCategoryStore(dbConn),
		TokenStore:    NewTokenStore(dbConn),
		CommentStore:  NewCommentStore(dbConn),
	}
}

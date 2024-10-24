package postgresql

import (
	"engine/internal/dto"
	"engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type CommentStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewCommentStore(db *pgxpool.Pool, log *zap.Logger) *CommentStore {
	return &CommentStore{
		db:  db,
		log: log,
	}
}

func (com *CommentStore) Create(commentDTO dto.Comment) (bool, error) {
	return true, nil
}

func (com *CommentStore) GetAllByPost() ([]model.Comment, error) {
	var comments []model.Comment

	return comments, nil
}

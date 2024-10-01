package postgresql

import (
	"engine/internal/dto"
	"engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type CommentStore struct {
	db *pgxpool.Pool
}

func NewCommentStore(db *pgxpool.Pool) *CommentStore {
	return &CommentStore{db: db}
}

func (com *CommentStore) Create(commentDTO dto.Comment) (bool, error) {
	return true, nil
}

func (com *CommentStore) GetAllByPost() ([]model.Comment, error) {
	var comments []model.Comment

	return comments, nil
}

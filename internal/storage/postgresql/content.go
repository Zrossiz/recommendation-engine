package postgresql

import (
	"context"
	"engine/internal/apperrors"
	"engine/internal/dto"
	"engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type ContentStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewContentStore(db *pgxpool.Pool, log *zap.Logger) *ContentStore {
	return &ContentStore{
		db:  db,
		log: log,
	}
}

func (co *ContentStore) Create(contentDTO dto.Content) (bool, error) {
	sql := `INSERT INTO content (category_id, title, description, link) VALUES ($1, $2, $3, $4)`
	_, err := co.db.Exec(
		context.Background(),
		sql,
		contentDTO.CategoryID,
		contentDTO.Title,
		contentDTO.Description,
		contentDTO.Link,
	)
	if err != nil {
		co.log.Error("create content error", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (co *ContentStore) Delete(id int64) (bool, error) {
	sql := `DELETE FROM content WHERE id = $1`
	_, err := co.db.Exec(context.Background(), sql, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (co *ContentStore) Update(contentID int64, contentDTO dto.Content) (bool, error) {
	sql := `UPDATE content SET category_id = $1, title = $2, description = $3, link = $4 WHERE id = $5`
	cmdTag, err := co.db.Exec(
		context.Background(),
		sql,
		contentDTO.CategoryID,
		contentDTO.Title,
		contentDTO.Description,
		contentDTO.Link,
		contentID,
	)

	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, apperrors.ErrNotFound
	}

	return true, nil
}

func (co *ContentStore) GetAllByCategory(categoryId int64, page int64) ([]model.Content, error) {
	sql := `SELECT category_id, title, description, link FROM content WHERE category_id = $1`
	rows, err := co.db.Query(context.Background(), sql, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var content []model.Content
	for rows.Next() {
		var cntnt model.Content
		err := rows.Scan(
			&cntnt.CategoryID,
			&cntnt.Title,
			&cntnt.Description,
			&cntnt.Link,
		)
		if err != nil {
			return nil, err
		}
		content = append(content, cntnt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

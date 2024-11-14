package postgresql

import (
	"context"
	"engine/internal/dto"
	"engine/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type UserInteractionsStore struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewUserInteractionsStore(db *pgxpool.Pool, log *zap.Logger) *UserInteractionsStore {
	return &UserInteractionsStore{
		db:  db,
		log: log,
	}
}

func (i *UserInteractionsStore) Create(dto dto.CreateInteraction) (bool, error) {
	sql := `INSERT INTO user_interactions (user_id, content_id, action, view_time) VALUES ($1, $2, $3, $4)`

	_, err := i.db.Exec(context.Background(), sql, dto.UserID, dto.ContentID, dto.Action, dto.ViewTime)
	if err != nil {
		i.log.Error("insert into user interactions", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (i *UserInteractionsStore) GetInteractionsByUser(userID int64) ([]dto.GetUserInteractions, error) {
	sql := `
		SELECT 
			u.id AS interaction_id,
			u.user_id,
			u.content_id,
			u.action,
			u.view_time,
			u.created_at AS interaction_created_at,
			ca.id AS category_id,
			ca.name AS category_name
		FROM user_interactions u
		LEFT JOIN content AS c ON u.content_id = c.id
		LEFT JOIN categories AS ca ON c.category_id = ca.id
		WHERE u.user_id = $1
		AND view_time >= 40
		ORDER BY u.created_at DESC
		LIMIT 100
	`

	rows, err := i.db.Query(context.Background(), sql, userID)
	if err != nil {
		i.log.Error("get interactions by user", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var resp []dto.GetUserInteractions
	for rows.Next() {
		var interaction dto.GetUserInteractions
		err := rows.Scan(
			&interaction.ID,
			&interaction.UserID,
			&interaction.ContentID,
			&interaction.Action,
			&interaction.ViewTime,
			&interaction.CreatedAt,
			&interaction.Category.ID,
			&interaction.Category.Name,
		)
		if err != nil {
			i.log.Error("scan user interaction", zap.Error(err))
			return nil, err
		}

		resp = append(resp, interaction)
	}

	if err = rows.Err(); err != nil {
		i.log.Error("rows error user interaction", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (i *UserInteractionsStore) GetCategoriesFromUserInteractions(userID int64) ([]model.Category, error) {
	sql := `
		WITH last_100_interactions AS (
			SELECT content_id
			FROM user_interactions
			WHERE user_id = $1
			AND view_time >= 40
			ORDER BY created_at DESC
			LIMIT 100
		)
		SELECT DISTINCT categories.id, categories.name
		FROM last_100_interactions
		JOIN content ON content.id = last_100_interactions.content_id
		JOIN categories ON categories.id = content.category_id
	`

	rows, err := i.db.Query(context.Background(), sql, userID)
	if err != nil {
		i.log.Error("GetCategoriesFromUserInteractions", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			i.log.Error("GetCategoriesFromUserInteractions scan", zap.Error(err))
			return nil, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		i.log.Error("rows error GetCategoriesFromUserInteractions", zap.Error(err))
		return nil, err
	}

	return categories, nil
}

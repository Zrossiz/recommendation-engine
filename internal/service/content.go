package service

import (
	"engine/internal/constants"
	"engine/internal/dto"
	"engine/internal/model"
	"math"

	"go.uber.org/zap"
)

type ContentStore interface {
	Create(contentDTO dto.Content) (bool, error)
	Delete(id int64) (bool, error)
	GetAllByCategory(categoryId int64, page int64) ([]model.Content, error)
	Update(contentID int64, contentDTO dto.Content) (bool, error)
	GetNewContentForUserFromCategories(userID int64, contentCategories map[string]int) ([]model.Content, error)
}

type ContentService struct {
	db             ContentStore
	interactionsDB UserInteractionsStore
	interestsDB    InterestsStore
	log            *zap.Logger
}

func NewContentService(
	db ContentStore,
	interactionsDB UserInteractionsStore,
	interestsDB InterestsStore,
	log *zap.Logger,
) *ContentService {
	return &ContentService{
		db:             db,
		interactionsDB: interactionsDB,
		interestsDB:    interestsDB,
		log:            log,
	}
}

func (c *ContentService) Create(contentDTO dto.Content) (bool, error) {
	_, err := c.db.Create(contentDTO)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *ContentService) GetRecommendations(userID int) ([]model.Content, error) {
	lastUserCategories, err := c.interactionsDB.GetCategoriesFromUserInteractions(int64(userID))
	if err != nil {
		return nil, err
	}

	lastUserInteractions, err := c.interactionsDB.GetInteractionsByUser(int64(userID))
	if err != nil {
		return nil, err
	}

	categoriesRating := make(map[string]float64)

	for _, category := range lastUserCategories {
		categoriesRating[category.Name] = 0
	}

	for _, interaction := range lastUserInteractions {
		points := math.Pow(float64(interaction.ViewTime)/100, 2) * 10

		switch interaction.Action {
		case constants.View:
			points += 1
		case constants.Comment:
			points += 2
		case constants.Like:
			points += 3
		case constants.Save:
			points += 4
		case constants.Repost:
			points += 5
		case constants.Dislike:
			points -= 100
		}
		categoriesRating[interaction.Category.Name] += points
	}

	userInterests, err := c.interestsDB.GetUserInterests(int64(userID))
	if err != nil {
		return nil, err
	}

	for _, interest := range userInterests {
		for _, interaction := range lastUserInteractions {
			if interest.CategoryID == interaction.Category.ID {
				categoriesRating[interaction.Category.Name] += 10
				break
			}
		}
	}

	contentCountForEveryCategory := make(map[string]int)
	for _, interaction := range lastUserInteractions {
		contentCountForEveryCategory[interaction.Category.Name]++
	}

	totalContent := len(lastUserInteractions)
	for category, count := range contentCountForEveryCategory {
		if count*100/totalContent > 50 {
			debuff := 0.3 * categoriesRating[category]
			categoriesRating[category] -= debuff
		}
	}

	var totalPoints float64
	for _, points := range categoriesRating {
		totalPoints += points
	}

	everyCategoryContentPercent := make(map[string]int)
	for key, value := range categoriesRating {
		percent := (value / totalPoints) * 100
		everyCategoryContentPercent[key] = int(percent)
	}

	everyCategoryContentQuantity := make(map[string]int)
	for category, percent := range everyCategoryContentPercent {
		everyCategoryContentQuantity[category] = percent * 17 / 100
	}

	generatedRecommedations, err := c.db.GetNewContentForUserFromCategories(int64(userID), everyCategoryContentQuantity)
	if err != nil {
		return nil, err
	}

	return generatedRecommedations, nil
}

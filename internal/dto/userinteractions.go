package dto

import "engine/internal/model"

type CreateInteraction struct {
	UserID    int64  `json:"user_id"`
	ContentID int64  `json:"content_id"`
	Action    string `json:"action"`    // like, dislike, view
	ViewTime  int16  `json:"view_time"` // percent watched of viewed content
}

type GetUserInteractions struct {
	model.Content
	Category model.Category
}

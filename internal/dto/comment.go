package dto

type Comment struct {
	UserID int64  `json:"user_id"`
	Text   string `json:"text"`
}

package dto

type Content struct {
	CategoryID  int64  `json:"category_id" validate:"required,min=1"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required,url"`
}

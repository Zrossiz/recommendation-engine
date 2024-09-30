package handler

type Handler struct {
	Category *CategoryHandler
	User     *UserHandler
	Content  *ContentHandler
}

func New() *Handler {
	return &Handler{
		Category: NewCategoryHandler(),
		User:     NewUserHandler(),
		Content:  NewContentHandler(),
	}
}

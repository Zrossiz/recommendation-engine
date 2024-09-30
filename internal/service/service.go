package service

type Service struct {
	Category *CategoryService
	User     *UserService
	Content  *ContentService
}

func New() *Service {
	return &Service{
		Category: NewCategoryService(),
		User:     NewUserService(),
		Content:  NewContentService(),
	}
}

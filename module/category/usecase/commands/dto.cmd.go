package categorycommands

type CreateCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

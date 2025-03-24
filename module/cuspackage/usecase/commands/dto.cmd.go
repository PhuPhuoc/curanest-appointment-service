package cuspackagecommands

type CreateCustomizedPackageDTO struct {
	Name string `json:"name" binding:"required"`
}
type CreateCustomizedTaskDTO struct {
	Name string `json:"name" binding:"required"`
}

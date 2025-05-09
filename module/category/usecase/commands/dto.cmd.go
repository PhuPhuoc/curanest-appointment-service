package categorycommands

import (
	"time"

	"github.com/google/uuid"
)

type CreateCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

type UpdateCategoryDTO struct {
	Id          uuid.UUID  `json:"id"`
	StaffId     *uuid.UUID `json:"staff-id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Thumbnail   string     `json:"thumbnail"`
	CreatedAt   *time.Time `json:"created_at"`
}

package categoryrepository

import (
	"time"

	"github.com/google/uuid"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

var (
	TABLE = `categories`

	FIELD = []string{
		"id",
		"staff_id",
		"name",
		"description",
	}

	UPDATE_FIELD = []string{
		"name",
		"description",
	}
)

type CategoryDTO struct {
	Id          uuid.UUID  `db:"id"`
	StaffId     *uuid.UUID `db:"staff_id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	CreatedAt   *time.Time `db:"created_at"`
}

func (dto *CategoryDTO) ToEntity() (*categorydomain.Category, error) {
	return categorydomain.NewCategory(
		dto.Id,
		dto.StaffId,
		dto.Name,
		dto.Description,
		dto.CreatedAt,
	)
}

func ToDTO(data *categorydomain.Category) *CategoryDTO {
	staffId := data.GetStaffId()
	return &CategoryDTO{
		Id:          data.GetID(),
		StaffId:     staffId,
		Name:        data.GetName(),
		Description: data.GetDescription(),
	}
}

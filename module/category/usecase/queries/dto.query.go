package categoryqueries

import (
	"github.com/google/uuid"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

type FilterCategoryDTO struct {
	Name string `json:"name"`
}

type CategoryDTO struct {
	Id          uuid.UUID  `json:"id"`
	StaffId     *uuid.UUID `json:"-"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Thumbnail   string     `json:"thumbnail"`
	StaffInfo   *StaffDTO  `json:"staff-info"`
}

func toDTO(data *categorydomain.Category) *CategoryDTO {
	dto := &CategoryDTO{
		Id:          data.GetID(),
		StaffId:     data.GetStaffId(),
		Name:        data.GetName(),
		Description: data.GetDescription(),
		Thumbnail:   data.GetThumbnail(),
	}
	return dto
}

type StaffDTO struct {
	NurseId      uuid.UUID `json:"nurse-id"`
	NursePicture string    `json:"nurse-picture"`
	NurseName    string    `json:"nurse-name"`
}

type StaffIdsQueryDTO struct {
	Ids []uuid.UUID `json:"ids"`
}

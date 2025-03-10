package servicequeries

import (
	"github.com/google/uuid"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

type FilterGetService struct {
	ServiceName string `json:"service-name"`
}

type ServiceDTO struct {
	Id          uuid.UUID `json:"id"`
	CategoryId  uuid.UUID `json:"category-id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EstDuration string    `json:"est-duration"`
	Status      string    `json:"status"`
}

func ToServiceDTO(entity *servicedomain.Service) *ServiceDTO {
	return &ServiceDTO{
		entity.GetID(),
		entity.GetCatetgoryID(),
		entity.GetName(),
		entity.GetDescription(),
		entity.GetEstDuration(),
		entity.GetStatus().String(),
	}
}

type CategoryDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
}

func ToCategoryDTO(entity *categorydomain.Category) *CategoryDTO {
	return &CategoryDTO{
		entity.GetID(),
		entity.GetName(),
		entity.GetDescription(),
		entity.GetThumbnail(),
	}
}

type ListServiceWithCategory struct {
	CategoryInfo CategoryDTO  `json:"category-info"`
	ListServices []ServiceDTO `json:"list-services"`
}

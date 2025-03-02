package servicequeries

import (
	"github.com/google/uuid"

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
	Thumbnail   string    `json:"thumbnail"`
	EstDuration string    `json:"est-duration"`
	Status      string    `json:"status"`
}

func ToServiceDTO(entity *servicedomain.Service) *ServiceDTO {
	return &ServiceDTO{
		entity.GetID(),
		entity.GetCatetgoryID(),
		entity.GetName(),
		entity.GetDescription(),
		entity.GetThumbnail(),
		entity.GetEstDuration(),
		entity.GetStatus().String(),
	}
}

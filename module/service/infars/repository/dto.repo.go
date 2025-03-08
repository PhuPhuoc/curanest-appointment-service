package servicerepository

import (
	"time"

	"github.com/google/uuid"

	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

var (
	TABLE = `services`

	CREATE_FIELD = []string{
		"id",
		"category_id",
		"name",
		"description",
		"est_duration",
	}

	GET_FIELD = []string{
		"id",
		"category_id",
		"name",
		"description",
		"est_duration",
		"status",
		"created_at",
	}

	UPDATE_FIELD = []string{
		"category_id",
		"name",
		"description",
		"est_duration",
	}
)

type ServiceDTO struct {
	Id          uuid.UUID  `db:"id"`
	CategoryId  uuid.UUID  `db:"category_id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	EstDuration string     `db:"est_duration"`
	Status      string     `db:"status"`
	CreatedAt   *time.Time `db:"created_at"`
}

func (dto *ServiceDTO) ToEntity() (*servicedomain.Service, error) {
	return servicedomain.NewService(
		dto.Id,
		dto.CategoryId,
		dto.Name,
		dto.Description,
		dto.EstDuration,
		servicedomain.Enum(dto.Status),
		dto.CreatedAt,
	)
}

func ToDTO(data *servicedomain.Service) *ServiceDTO {
	return &ServiceDTO{
		Id:          data.GetID(),
		CategoryId:  data.GetCatetgoryID(),
		Name:        data.GetName(),
		Description: data.GetDescription(),
		EstDuration: data.GetEstDuration(),
		Status:      data.GetStatus().String(),
	}
}

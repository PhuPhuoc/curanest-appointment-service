package categoryqueries

import (
	"context"

	"github.com/google/uuid"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

type Queries struct {
	GetAllCategories *getCategoriesHandler

	FindCategoryById *findCategoryByIdHandler
}

type Builder interface {
	BuildCategoryQueryRepo() CategoryQueryRepo
	BuildExternalNursingServiceInQuery() ExternalNursingService
}

func NewCategoryQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetAllCategories: NewGetCategoriesHandler(
			b.BuildCategoryQueryRepo(),
			b.BuildExternalNursingServiceInQuery(),
		),

		FindCategoryById: NewFindCategoryByIdHandler(
			b.BuildCategoryQueryRepo(),
		),
	}
}

type CategoryQueryRepo interface {
	GetCategories(ctx context.Context, filter *FilterCategoryDTO) ([]categorydomain.Category, error)
	FindCategoryById(ctx context.Context, cateId uuid.UUID) (*categorydomain.Category, error)
}

type ExternalNursingService interface {
	GetStaffsRPC(ctx context.Context, ids *StaffIdsQueryDTO) ([]StaffDTO, error)
}

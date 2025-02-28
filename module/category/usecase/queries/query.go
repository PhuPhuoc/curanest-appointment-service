package categoryqueries

import (
	"context"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

type Queries struct {
	GetAllCategories *getCategoriesHandler
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
	}
}

type CategoryQueryRepo interface {
	GetCategories(ctx context.Context, filter *FilterCategoryDTO) ([]categorydomain.Category, error)
}

type ExternalNursingService interface {
	GetStaffsRPC(ctx context.Context, ids *StaffIdsQueryDTO) ([]StaffDTO, error)
}

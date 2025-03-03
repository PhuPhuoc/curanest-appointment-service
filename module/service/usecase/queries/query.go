package servicequeries

import (
	"context"

	"github.com/google/uuid"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
	categoryqueries "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/queries"
	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

type Queries struct {
	GetByCategory      *getServicesByCategoryHandler
	GetGroupByCategory *getServicesGroupByCategoryHandler
}

type Builder interface {
	BuildServiceQueryRepo() ServiceQueryRepo
	BuildCategoryFetcher() CategoryFetcher
}

func NewServiceQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetByCategory: NewGetServicesByCategoryHandler(
			b.BuildServiceQueryRepo(),
		),
		GetGroupByCategory: NewGetServicesGroupByCategoryHandler(
			b.BuildServiceQueryRepo(),
			b.BuildCategoryFetcher(),
		),
	}
}

type ServiceQueryRepo interface {
	GetServicesByCategoryAndFilter(ctx context.Context, cateId uuid.UUID, filter FilterGetService) ([]servicedomain.Service, error)
	// GetServicesWithFilter(ctx context.Context, filter FilterGetService) ([]servicedomain.Service, error)
}

type CategoryFetcher interface {
	GetCategories(ctx context.Context, filter *categoryqueries.FilterCategoryDTO) ([]categorydomain.Category, error)
}

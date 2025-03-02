package servicerepository

import (
	"context"

	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
)

func (repo *serviceRepo) GetServicesWithFilter(ctx context.Context, filter servicequeries.FilterGetService) ([]servicedomain.Service, error) {
	return nil, nil
}

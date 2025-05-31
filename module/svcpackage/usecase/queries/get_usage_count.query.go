package svcpackagequeries

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getSvcPackagesUsageCountHandler struct {
	queryRepo SvcPackageQueryRepo
}

func NewGetServicePackagesUsageCountHandler(queryRepo SvcPackageQueryRepo) *getSvcPackagesUsageCountHandler {
	return &getSvcPackagesUsageCountHandler{
		queryRepo: queryRepo,
	}
}

func (h *getSvcPackagesUsageCountHandler) Handle(ctx context.Context, categoryId uuid.UUID) ([]SvcPackageUsageCountDTO, error) {
	entities, err := h.queryRepo.GetSvcPackageUsageCount(ctx, categoryId)
	if err != nil {
		return []SvcPackageUsageCountDTO{}, common.NewInternalServerError().
			WithReason("cannot get list service-package-usage-count").
			WithInner(err.Error())
	}
	if len(entities) == 0 {
		return []SvcPackageUsageCountDTO{}, nil
	}

	dtos := make([]SvcPackageUsageCountDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = *toSvcPackageUsageCountDTO(&entity)
	}

	return dtos, nil
}

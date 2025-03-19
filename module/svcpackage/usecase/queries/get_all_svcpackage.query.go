package svcpackagequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type getSvcPackagesHandler struct {
	queryRepo SvcPackageQueryRepo
}

func NewGetServicePackagesHandler(queryRepo SvcPackageQueryRepo) *getSvcPackagesHandler {
	return &getSvcPackagesHandler{
		queryRepo: queryRepo,
	}
}

func (h *getSvcPackagesHandler) Handle(ctx context.Context, serviceId uuid.UUID) ([]ServicePackageDTO, error) {
	entities, err := h.queryRepo.GetSvcPackges(ctx, serviceId)
	if err != nil {
		return []ServicePackageDTO{}, common.NewInternalServerError().
			WithReason("cannot get list service-package by service-id '" + serviceId.String() + "'").
			WithInner(err.Error())
	}
	if len(entities) == 0 {
		return []ServicePackageDTO{}, nil
	}

	dtos := make([]ServicePackageDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = *toServicePackageDTO(&entity)
	}

	return dtos, nil
}

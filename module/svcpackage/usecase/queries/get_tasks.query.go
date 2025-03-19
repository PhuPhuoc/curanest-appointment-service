package svcpackagequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type getSvcTasksHandler struct {
	queryRepo SvcPackageQueryRepo
}

func NewGetServiceTasksHandler(queryRepo SvcPackageQueryRepo) *getSvcTasksHandler {
	return &getSvcTasksHandler{
		queryRepo: queryRepo,
	}
}

func (h *getSvcTasksHandler) Handle(ctx context.Context, svcpackageId uuid.UUID) ([]ServiceTaskDTO, error) {
	entities, err := h.queryRepo.GetSvcTasks(ctx, svcpackageId)
	if err != nil {
		return []ServiceTaskDTO{}, common.NewInternalServerError().
			WithReason("cannot get list service-package by service-package-id '" + svcpackageId.String() + "'").
			WithInner(err.Error())
	}
	if len(entities) == 0 {
		return []ServiceTaskDTO{}, nil
	}

	dtos := make([]ServiceTaskDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = *toServiceTaskDTO(&entity)
	}

	return dtos, nil
}

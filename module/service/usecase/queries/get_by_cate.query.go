package servicequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type getServicesByCategoryHandler struct {
	queryRepo ServiceQueryRepo
}

func NewGetServicesByCategoryHandler(queryRepo ServiceQueryRepo) *getServicesByCategoryHandler {
	return &getServicesByCategoryHandler{
		queryRepo: queryRepo,
	}
}

func (h *getServicesByCategoryHandler) Handle(ctx context.Context, cateId uuid.UUID, filter FilterGetService) ([]ServiceDTO, error) {
	entities, err := h.queryRepo.GetServicesByCategoryAndFilter(ctx, cateId, filter)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list services by category-id '" + cateId.String() + "'").
			WithInner(err.Error())
	}

	dtos := make([]ServiceDTO, len(entities))

	for i := range entities {
		dto := ToServiceDTO(&entities[i])
		dtos[i] = *dto
	}

	return dtos, nil
}

package cuspackagequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type findCusPackageTaskHandler struct {
	queryRepo CusPackageQueryRepo
}

func NewFindCusPackageTaskDetailHandler(queryRepo CusPackageQueryRepo) *findCusPackageTaskHandler {
	return &findCusPackageTaskHandler{
		queryRepo: queryRepo,
	}
}

func (h *findCusPackageTaskHandler) Handle(ctx context.Context, filter *FilterGetCusPackageTaskDTO) (PackageTaskResponse, error) {
	cusPackage, err := h.queryRepo.FindCusPackage(ctx, filter.CusPackageId)
	if err != nil {
		return PackageTaskResponse{}, common.NewInternalServerError().
			WithReason("cannot find customized package").
			WithInner(err.Error())
	}

	cusTasks, err := h.queryRepo.FindCusTasks(ctx, filter.CusPackageId, filter.EstDate)
	if err != nil {
		return PackageTaskResponse{}, common.NewInternalServerError().
			WithReason("cannot find customized tasks").
			WithInner(err.Error())
	}
	if len(cusTasks) == 0 {
		return PackageTaskResponse{}, nil
	}

	dtos := make([]CusTaskDTO, len(cusTasks))
	for i, entity := range cusTasks {
		dtos[i] = *toCusTaskDTO(&entity)
	}

	response := PackageTaskResponse{
		Package: toCusPackageDTO(cusPackage),
		Tasks:   dtos,
	}

	return response, nil
}

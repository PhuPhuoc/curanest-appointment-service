package servicequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getServicesGroupByCategoryHandler struct {
	queryRepo   ServiceQueryRepo
	cateFetcher CategoryFetcher
}

func NewGetServicesGroupByCategoryHandler(queryRepo ServiceQueryRepo, cateFetcher CategoryFetcher) *getServicesGroupByCategoryHandler {
	return &getServicesGroupByCategoryHandler{
		queryRepo:   queryRepo,
		cateFetcher: cateFetcher,
	}
}

func (h *getServicesGroupByCategoryHandler) Handle(ctx context.Context, filter FilterGetService) ([]ListServiceWithCategory, error) {
	list_cate, err := h.cateFetcher.GetCategories(ctx, nil)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list category").
			WithInner(err.Error())
	}

	dtos := make([]ListServiceWithCategory, len(list_cate))
	for i := range list_cate {
		dtos[i].CategoryInfo = *ToCategoryDTO(&list_cate[i])
	}

	for i := range dtos {
		cateId := dtos[i].CategoryInfo.Id
		list_service, err := h.queryRepo.GetServicesByCategoryAndFilter(
			ctx,
			cateId,
			filter,
		)
		if err != nil {
			return nil, common.NewInternalServerError().
				WithReason("cannot get list service of categogy - (id: " + cateId.String()).
				WithInner(err.Error())
		}

		servicedtos := make([]ServiceDTO, len(list_service))
		for i := range list_service {
			servicedtos[i] = *ToServiceDTO(&list_service[i])
		}

		dtos[i].ListServices = servicedtos
	}

	return dtos, nil
}

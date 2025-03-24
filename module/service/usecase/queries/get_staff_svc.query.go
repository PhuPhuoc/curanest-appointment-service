package servicequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getStaffServiceHandler struct {
	queryRepo ServiceQueryRepo
	cateFetch CategoryFetcher
}

func NewGetStaffServicesHandler(queryRepo ServiceQueryRepo, cateFetch CategoryFetcher) *getStaffServiceHandler {
	return &getStaffServiceHandler{
		queryRepo: queryRepo,
		cateFetch: cateFetch,
	}
}

func (h *getStaffServiceHandler) Handle(ctx context.Context, filter FilterGetService) (*ListServiceWithCategory, error) {
	requester, ok := ctx.Value(common.KeyRequester).(common.Requester)
	if !ok {
		return nil, common.NewUnauthorizedError()
	}
	userRole := requester.Role()
	if userRole != "staff" {
		if !ok {
			return nil, common.NewUnauthorizedError().WithReason("your role is not staff")
		}
	}
	staffId := requester.UserId()

	cate, err := h.cateFetch.GetCategoryOfStaff(ctx, staffId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get category data of this staff account").
			WithInner(err.Error())
	}

	entities, err := h.queryRepo.GetServicesByCategoryAndFilter(ctx, cate.GetID(), filter)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list services by category-id '" + cate.GetID().String() + "'").
			WithInner(err.Error())
	}

	dtos := make([]ServiceDTO, len(entities))

	for i := range entities {
		dto := ToServiceDTO(&entities[i])
		dtos[i] = *dto
	}

	result := &ListServiceWithCategory{
		CategoryInfo: *ToCategoryDTO(cate),
		ListServices: dtos,
	}

	return result, nil
}

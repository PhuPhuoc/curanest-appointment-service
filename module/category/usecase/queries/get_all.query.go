package categoryqueries

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getCategoriesHandler struct {
	queryRepo  CategoryQueryRepo
	nursingRPC ExternalNursingService
}

func NewGetCategoriesHandler(queryRepo CategoryQueryRepo, nursingRPC ExternalNursingService) *getCategoriesHandler {
	return &getCategoriesHandler{
		queryRepo:  queryRepo,
		nursingRPC: nursingRPC,
	}
}

func (h *getCategoriesHandler) Handle(ctx context.Context, filter *FilterCategoryDTO) ([]CategoryDTO, error) {
	entities, err := h.queryRepo.GetCategories(ctx, filter)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("error at GetCategories-repo").
			WithInner(err.Error())
	}

	list_dto := make([]CategoryDTO, len(entities))
	list_ids := make([]uuid.UUID, len(entities))

	for i := range entities {
		dto := toDTO(&entities[i])
		list_dto[i] = *dto
		if dto.StaffId != nil {
			list_ids[i] = *dto.StaffId
		}
	}

	staffQueryDTO := StaffIdsQueryDTO{
		Ids: list_ids,
	}

	if len(list_dto) > 0 {
		staffs, err := h.nursingRPC.GetStaffsRPC(ctx, &staffQueryDTO)
		if err != nil {
			return nil, err
		}

		for i := range list_dto {
			dtoStaffId := list_dto[i].StaffId
			if dtoStaffId == nil {
				continue
			}
			for j := range staffs {
				staffId := staffs[j].NurseId
				fmt.Printf("dto_staff_id: %v ___ staffs_id: %v \n", *dtoStaffId, staffId)
				if *dtoStaffId == staffId {
					list_dto[i].StaffInfo = &staffs[j]
					break
				}
			}
		}
	}

	return list_dto, nil
}

package categorycommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type addStaffToCategoryHandler struct {
	cmdRepo  CategoryCommandRepo
	nurseRPC ExternalAccountService
}

func NewAddStaffToCategoryHandler(cmdRepo CategoryCommandRepo, nurseRPC ExternalAccountService) *addStaffToCategoryHandler {
	return &addStaffToCategoryHandler{
		cmdRepo:  cmdRepo,
		nurseRPC: nurseRPC,
	}
}

func (h *addStaffToCategoryHandler) Handle(ctx context.Context, cateId, staffId uuid.UUID) error {
	if err := h.cmdRepo.AddStaffToCategory(ctx, cateId, staffId); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot apply new staff for this category - " + cateId.String()).
			WithInner(err.Error())
	}

	if err := h.nurseRPC.UpdateAccountRoleRPC(ctx, staffId, "staff"); err != nil {
		_ = h.cmdRepo.RemoveStaffOfCategory(ctx, cateId)
		return err
	}

	return nil
}

package categorycommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type removeStaffToCategoryHandler struct {
	cmdRepo  CategoryCommandRepo
	nurseRPC ExternalAccountService
}

func NewRemoveStaffToCategoryHandler(cmdRepo CategoryCommandRepo, nurseRPC ExternalAccountService) *removeStaffToCategoryHandler {
	return &removeStaffToCategoryHandler{
		cmdRepo:  cmdRepo,
		nurseRPC: nurseRPC,
	}
}

func (h *removeStaffToCategoryHandler) Handle(ctx context.Context, cateId, staffId uuid.UUID) error {
	if err := h.cmdRepo.RemoveStaffOfCategory(ctx, cateId); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot remove staff in this category - " + cateId.String()).
			WithInner(err.Error())
	}

	if err := h.nurseRPC.UpdateAccountRoleRPC(ctx, staffId, "nurse"); err != nil {
		_ = h.cmdRepo.AddStaffToCategory(ctx, cateId, staffId)
		return err
	}

	return nil
}

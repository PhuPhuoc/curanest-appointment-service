package categoryhttpservice

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

// @Summary		add staff to category (admin)
// @Description	add staff to category (admin)
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			category-id	path		string					true	"category ID (UUID)"
// @Param			staff-id	path		string					true	"staff ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/categories/{category-id}/staff/{staff-id} [patch]
// @Security		ApiKeyAuth
func (s *categoryHttpService) handleAddStaffForCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cateId := ctx.Param("category-id")
		if cateId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing category-id"))
			return
		}

		cateUUID, err := uuid.Parse(cateId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("category-id is invalid (not a uuid)"))
			return
		}

		staffId := ctx.Param("staff-id")
		if staffId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing staff-id"))
			return
		}

		staffUUID, err := uuid.Parse(staffId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("staff-id is invalid (not a uuid)"))
			return
		}

		cateEntity, err := s.query.FindCategoryById.Handle(ctx, cateUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if cateEntity.StaffId != nil && *cateEntity.StaffId == staffUUID {
			messErr := fmt.Errorf("staff(id: %v) is already a staff of this category", cateUUID)
			common.ResponseError(ctx, common.NewBadRequestError().WithReason(messErr.Error()))
			return
		}

		if cateEntity.StaffId != nil {
			messErr := fmt.Errorf("staff already exists for category '%v'", cateEntity.Name)
			common.ResponseError(ctx, common.NewBadRequestError().WithReason(messErr.Error()))
			return
		}

		if err := s.command.AddStaff.Handle(ctx.Request.Context(), cateUUID, staffUUID); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

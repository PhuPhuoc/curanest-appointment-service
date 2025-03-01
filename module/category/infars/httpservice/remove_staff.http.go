package categoryhttpservice

import (
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		remove staff to category (admin)
// @Description	remove staff to category (admin)
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			category-id	path		string					true	"category ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/categories/{category-id}/staff/remove [patch]
// @Security		ApiKeyAuth
func (s *categoryHttpService) handleRemoveStaffForCategory() gin.HandlerFunc {
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

		cateEntity, err := s.query.FindCategoryById.Handle(ctx, cateUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if cateEntity.StaffId == nil {
			messErr := fmt.Errorf("there is currently no staff managing the category '%v'", cateEntity.Name)
			common.ResponseError(ctx, common.NewBadRequestError().WithReason(messErr.Error()))
			return
		}

		if err := s.command.RemoveStaff.Handle(ctx.Request.Context(), cateUUID, *cateEntity.StaffId); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

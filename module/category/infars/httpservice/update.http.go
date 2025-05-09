package categoryhttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorycommands "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/commands"
	"github.com/gin-gonic/gin"
)

// @Summary		update category
// @Description	update category
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			create	form		body					categorycommands.UpdateCategoryDTO	true	"category update data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/categories [put]
// @Security		ApiKeyAuth
func (s *categoryHttpService) handleUpdateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto categorycommands.UpdateCategoryDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.UpdateCategory.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}

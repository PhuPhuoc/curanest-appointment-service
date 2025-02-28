package categoryhttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorycommands "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/commands"
	"github.com/gin-gonic/gin"
)

//	@Summary		create new category
//	@Description	create new category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			create	form		body					categorycommands.CreateCategoryDTO	true	"account creation data"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/api/v1/categories [post]
//	@Security		ApiKeyAuth
func (s *categoryHttpService) handleCreateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto categorycommands.CreateCategoryDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreateCategory.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}

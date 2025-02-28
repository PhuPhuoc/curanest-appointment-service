package categoryhttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categoryqueries "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/queries"
)

// @Summary		get categories
// @Description	get categories by name
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			name	query		string					false	"category name"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/categories [get]
// @Security		ApiKeyAuth
func (s *categoryHttpService) handleGetCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")

		dto := categoryqueries.FilterCategoryDTO{
			Name: name,
		}
		dtos, err := s.query.GetAllCategories.Handle(ctx.Request.Context(), &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, dtos)
	}
}

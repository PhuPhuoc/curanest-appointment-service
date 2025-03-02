package servicehttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
)

// @Summary		get list service by category-id (admin)
// @Description	get list service by category-id (admin)
// @Tags			services
// @Accept			json
// @Produce		json
// @Param			category-id		path		string					true	"category ID (UUID)"
// @Param			service-name	query		string					false	"services name"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/categories/{category-id}/services [get]
// @Security		ApiKeyAuth
func (s *serviceHttpService) handleGetServiceByCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cateId := ctx.Param("category-id")
		if cateId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing category-id"))
			return
		}
		cateUUID, err := uuid.Parse(cateId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("category-id invalid (not a uuid)"))
			return
		}

		// add more field for fitler here
		serviceName := ctx.Query("service-name")
		filter := servicequeries.FilterGetService{
			ServiceName: serviceName,
		}

		data, err := s.query.GetByCategory.Handle(ctx.Request.Context(), cateUUID, filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, data)
	}
}

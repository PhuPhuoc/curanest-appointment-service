package servicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
	"github.com/gin-gonic/gin"
)

//	@Summary		get list service with category (guest)
//	@Description	get list service with category (guest)
//	@Tags			services
//	@Accept			json
//	@Param			service-name	query	string	false	"services name"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"data"
//	@Failure		400	{object}	error					"Bad request error"
//	@Router			/api/v1/services/group-by-category [get]
func (s *serviceHttpService) handleGetServiceGroupByCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		serviceName := ctx.Query("service-name")
		filter := servicequeries.FilterGetService{
			ServiceName: serviceName,
		}

		data, err := s.query.GetGroupByCategory.Handle(ctx.Request.Context(), filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, data)
	}
}

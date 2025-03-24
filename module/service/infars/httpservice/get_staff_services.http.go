package servicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		get list service of staff (staff)
// @Description	get list service of staff (staff)
// @Tags			services
// @Accept			json
// @Produce		json
// @Param			service-name	query		string					false	"services name"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/staff/services [get]
// @Security		ApiKeyAuth
func (s *serviceHttpService) handleGetServiceOfStaff() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// add more field for fitler here
		serviceName := ctx.Query("service-name")
		filter := servicequeries.FilterGetService{
			ServiceName: serviceName,
		}

		data, err := s.query.GetStaffServices.Handle(ctx.Request.Context(), filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, data)
	}
}

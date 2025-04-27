package appointmenthttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		verify nurses and dates
// @Description	verify nurses and dates are ready for new appointment
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			create	form		body					appointmentqueries.CheckNursesAvailabilityRequestDTO	true	"nurses and dates mapping"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/appointments/verify-nurses-dates [post]
// @Security		ApiKeyAuth
func (s *appointmentHttpService) handleVerifyNurseWithDate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dto := &appointmentqueries.CheckNursesAvailabilityRequestDTO{}
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		response, err := s.query.CheckNursesAvailability.Handle(ctx, dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, response)
	}
}

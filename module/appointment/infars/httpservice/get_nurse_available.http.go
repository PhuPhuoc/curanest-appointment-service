package appointmenthttpservice

import (
	"strconv"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get nursing available
// @Description	get nursing available with service of appointment and date
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			service-id		query		string					true	"service ID (UUID)"
// @Param			est-date		query		string					true	"est date (YYYY-MM-DDTHH:MM:SSZ, e.g., 2025-05-16T01:00:00Z)"
// @Param			est-duration	query		int						true	"est duration"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/appointments/nursing-available [get]
// @Security		ApiKeyAuth
func (s *appointmentHttpService) handleGetNursingAvailable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var serviceUUID uuid.UUID
		var estDate time.Time
		var estDuration int
		var err error
		if serviceId := ctx.Query("service-id"); serviceId != "" {
			serviceUUID, err = uuid.Parse(serviceId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a UUID)"))
				return
			}
		}

		if estDateStr := ctx.Query("est-date"); estDateStr != "" {
			estDate, err = time.Parse(time.RFC3339, estDateStr)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("est-date invalid (use YYYY-MM-DDTHH:MM:SSZ, e.g., 2025-05-16T01:00:00Z)"))
				return
			}
		}

		if estDurationStr := ctx.Query("est-duration"); estDurationStr != "" {
			estDuration, err = strconv.Atoi(estDurationStr)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("est_duration invalid (must be an integer)"))
				return
			}
		}

		nurses, err := s.query.GetNursingAvailable.Handle(ctx, serviceUUID, estDate, estDuration)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, nurses)
	}
}

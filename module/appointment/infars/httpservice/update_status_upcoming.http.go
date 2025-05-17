package appointmenthttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		update status of appointment to upcoming
// @Description	update status of appointment to upcoming
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			appointment-id	path		string					true	"appointment ID (UUID)"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/appointments/{appointment-id}/update-status-upcoming [patch]
// @Security		ApiKeyAuth
func (s *appointmentHttpService) handleUpdateStatusUpcoming() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var appointmentUUID uuid.UUID
		var err error
		if appointmentId := ctx.Param("appointment-id"); appointmentId != "" {
			appointmentUUID, err = uuid.Parse(appointmentId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a UUID)"))
				return
			}
		}

		appointmentDTO, err := s.query.FindAppointmentById.Handle(ctx.Request.Context(), appointmentUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		appointmentEntity, _ := appointmentDTO.ToAppointmentDomain()

		if err := s.command.UpdateStatusUpcoming.Handle(ctx.Request.Context(), appointmentEntity); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

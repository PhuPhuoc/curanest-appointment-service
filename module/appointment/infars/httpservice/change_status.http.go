package appointmenthttpservice

import (
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get appointment by filter option
// @Description	get appointment by filter option
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			appointment-id	path		string					true	"appointment ID (UUID)"
// @Param			new-status		query		string					true	"new status to update appointment's status"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/appointments/{appointment-id}/status [patch]
// @Security		ApiKeyAuth
func (s *appointmentHttpService) handleChangeAppointmentStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var appointmentUUID uuid.UUID
		var newStatus appointmentdomain.AppointmentStatus
		var err error
		if appointmentId := ctx.Param("appointment-id"); appointmentId != "" {
			appointmentUUID, err = uuid.Parse(appointmentId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a UUID)"))
				return
			}
		}

		if status := ctx.Query("new-status"); status != "" {
			newStatus = appointmentdomain.EnumAppointmentStatus(status)
			if newStatus == appointmentdomain.AppStatusUnknow {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason(fmt.Sprintf("new-status: %v invalid", status)))
				return
			}
		}

		appointmentDTO, err := s.query.FindAppointmentById.Handle(ctx.Request.Context(), appointmentUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		appointmentEntity, _ := appointmentDTO.ToAppointmentDomain()

		if err := s.command.UpdateApointmentStatus.Handle(ctx.Request.Context(), newStatus, appointmentEntity); err != nil {

			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

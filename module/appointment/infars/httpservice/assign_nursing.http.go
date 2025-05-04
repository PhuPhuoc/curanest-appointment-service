package appointmenthttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		assign nursing for appointment (staff)
// @Description	assign nursing for appointment (staff)
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			appointment-id	path		string					true	"appointment ID (UUID)"
// @Param			nursing-id		path		string					true	"nursing ID (UUID)"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/appointments/{appointment-id}/assign-nursing/{nursing-id} [patch]
// @Security		ApiKeyAuth
func (s *appointmentHttpService) handleAssignNurseToAppointment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var appointmentUUID uuid.UUID
		var err error
		if appointmentId := ctx.Param("appointment-id"); appointmentId != "" {
			appointmentUUID, err = uuid.Parse(appointmentId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("appointment-id invalid (not a UUID)"))
				return
			}
		}

		var nursingUUID uuid.UUID
		if nursingId := ctx.Param("nursing-id"); nursingId != "" {
			nursingUUID, err = uuid.Parse(nursingId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("nursing-id invalid (not a UUID)"))
				return
			}
		}

		appointmentDTO, err := s.query.FindAppointmentById.Handle(ctx.Request.Context(), appointmentUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		appointmentEntity, _ := appointmentDTO.ToAppointmentDomain()

		if err := s.command.AssigneNursing.Handle(ctx, &nursingUUID, appointmentEntity); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

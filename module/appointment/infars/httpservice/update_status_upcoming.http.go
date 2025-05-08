package appointmenthttpservice

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
// @Param			origin-code		query		string					false	"origin code (current location of nursing - lat/lng"
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

		var originCode *string
		if origin := ctx.Query("origin-code"); origin != "" {
			originCode = &origin
		}

		if !isValidLatLng(*originCode) {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason(fmt.Sprintf("origin-code: %v invalid (must be lat,lng)", originCode)))
			return
		}

		appointmentDTO, err := s.query.FindAppointmentById.Handle(ctx.Request.Context(), appointmentUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		appointmentEntity, _ := appointmentDTO.ToAppointmentDomain()

		if err := s.command.UpdateStatusUpcoming.Handle(ctx.Request.Context(), originCode, appointmentEntity); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

func isValidLatLng(originCode string) bool {
	pattern := `^[-]?[0-9]*\.?[0-9]+,[-]?[0-9]*\.?[0-9]+$`
	matched, err := regexp.MatchString(pattern, originCode)
	if err != nil || !matched {
		return false
	}

	// Tách lat và lng
	parts := strings.Split(originCode, ",")
	if len(parts) != 2 {
		return false
	}

	lat, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return false
	}
	lng, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return false
	}

	// Kiểm tra phạm vi hợp lệ
	if lat < -90 || lat > 90 {
		return false
	}
	if lng < -180 || lng > 180 {
		return false
	}

	return true
}

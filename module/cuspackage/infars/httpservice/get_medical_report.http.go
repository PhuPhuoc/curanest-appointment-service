package cuspackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get medical record
// @Description	get medical record
// @Tags			medical reports
// @Accept			json
// @Produce		json
// @Param			appointment-id	path		string					true	"custask ID (UUID)"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/medical-record/{appointment-id} [get]
// @Security		ApiKeyAuth
func (s *cusPackageHttpService) handleGetMedicalRecord() gin.HandlerFunc {
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

		medicalRecordDTO, err := s.query.FindMedicalRecordByAppsId.Handle(ctx.Request.Context(), appointmentUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, medicalRecordDTO)
	}
}

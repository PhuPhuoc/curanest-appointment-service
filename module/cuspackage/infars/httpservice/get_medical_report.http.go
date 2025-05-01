package cuspackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		get medical record
//	@Description	get medical record
//	@Tags			medical reports
//	@Accept			json
//	@Produce		json
//	@Param			medical-record-id	path		string					true	"custask ID (UUID)"
//	@Success		200					{object}	map[string]interface{}	"data"
//	@Failure		400					{object}	error					"Bad request error"
//	@Router			/api/v1/medical-record/{medical-record-id} [get]
//	@Security		ApiKeyAuth
func (s *cusPackageHttpService) handleGetMedicalRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var medicalRecordUUID uuid.UUID
		var err error
		if medicalRecordId := ctx.Param("medical-record-id"); medicalRecordId != "" {
			medicalRecordUUID, err = uuid.Parse(medicalRecordId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("medical-record-id invalid (not a UUID)"))
				return
			}
		}

		medicalRecordDTO, err := s.query.FindMedicalRecordById.Handle(ctx.Request.Context(), medicalRecordUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, medicalRecordDTO)
	}
}

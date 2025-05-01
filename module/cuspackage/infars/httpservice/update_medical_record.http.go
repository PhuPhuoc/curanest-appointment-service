package cuspackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		update medical record
// @Description	update medical record
// @Tags			medical reports
// @Accept			json
// @Produce		json
// @Param			medical-record-id	path		string					true										"medical-record ID (UUID)"
// @Param			update				form		body					cuspackagecommands.UpdateMedicalRecordDTO	true	"customized package and task creation data"
// @Success		200					{object}	map[string]interface{}	"data"
// @Failure		400					{object}	error					"Bad request error"
// @Router			/api/v1/medical-record/{medical-record-id} [patch]
// @Security		ApiKeyAuth
func (s *cusPackageHttpService) handleUpdateMedicalRecord() gin.HandlerFunc {
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

		var dto cuspackagecommands.UpdateMedicalRecordDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		medicalRecordDTO, err := s.query.FindMedicalRecordById.Handle(ctx.Request.Context(), medicalRecordUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		medicalRecordEntity, _ := medicalRecordDTO.ToMedicalRecordEntity()

		if err := s.cmd.UpdateMedicalRecord.Handle(ctx.Request.Context(), dto, medicalRecordEntity); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

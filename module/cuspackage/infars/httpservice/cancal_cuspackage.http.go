package cuspackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		cancel customized package (client)
//	@Description	create customized package (client)
//	@Tags			customized packages
//	@Accept			json
//	@Produce		json
//	@Param			cuspackage-id	path		string					true	"custask ID (UUID)"
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	error					"Bad request error"
//	@Router			/api/v1/cuspackage/{cuspackage-id}/cancel [patch]
//	@Security		ApiKeyAuth
func (s *cusPackageHttpService) handleCancelCustomizedPackage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cuspackageUUID uuid.UUID
		var err error
		if cuspackageId := ctx.Param("cuspackage-id"); cuspackageId != "" {
			cuspackageUUID, err = uuid.Parse(cuspackageId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("custask-id invalid (not a UUID)"))
				return
			}
		}

		cuspackageDTO, err := s.query.FindCuspackageById.Handle(ctx, cuspackageUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		cuspackageEntity, _ := cuspackageDTO.ToCusPackageEntity()

		err = s.cmd.CancelPackage.Handle(ctx.Request.Context(), cuspackageEntity)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

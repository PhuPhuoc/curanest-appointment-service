package cuspackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		change custask status to done
//	@Description	change custask status to done
//	@Tags			customized packages
//	@Accept			json
//	@Produce		json
//	@Param			custask-id	path		string					true	"custask ID (UUID)"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/cuspackage/custask/{custask-id}/update-status-done [patch]
//	@Security		ApiKeyAuth
func (s *cusPackageHttpService) handleUpdateCustaskStatusDone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var custaskUUID uuid.UUID
		var err error
		if custaskId := ctx.Param("custask-id"); custaskId != "" {
			custaskUUID, err = uuid.Parse(custaskId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("custask-id invalid (not a UUID)"))
				return
			}
		}

		custaskDTO, err := s.query.FindCustaskById.Handle(ctx.Request.Context(), custaskUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		custaskEntity, _ := custaskDTO.ToCusTaskEntity()

		if err := s.cmd.UpdateCustaskStatusDone.Handle(ctx, custaskEntity); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

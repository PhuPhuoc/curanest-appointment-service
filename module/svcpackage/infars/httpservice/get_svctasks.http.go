package svcpackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		get list service-tasks by svcpackage-id
//	@Description	get list service-tasks by svcpackage-id
//	@Tags			service packages
//	@Accept			json
//	@Produce		json
//	@Param			svcpackage-id	path		string					true	"service package ID (UUID)"
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	error					"Bad request error"
//	@Router			/api/v1/svcpackage/{svcpackage-id}/svctask [get]
func (s *svcPackageHttpService) handleGetServicTasks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		svcpackageId := ctx.Param("svcpackage-id")
		if svcpackageId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing svcpackage-id"))
			return
		}
		svcpackageUUID, err := uuid.Parse(svcpackageId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("svcpackage-id invalid (not a uuid)"))
			return
		}

		data, err := s.query.GetServiceTasks.Handle(ctx.Request.Context(), svcpackageUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, data)
	}
}

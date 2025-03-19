package svcpackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		update service task (staff)
// @Description	update service task (staff)
// @Tags			service packages
// @Accept			json
// @Produce		json
// @Param			svcpackage-id	path		string					true									"category ID (UUID)"
// @Param			svctask-id		path		string					true									"service-task ID (UUID)"
// @Param			create			form		body					svcpackagecommands.UpdateServiceTaskDTO	true	"service task update information"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/svcpackage/{svcpackage-id}/svctask/{svctask-id} [put]
// @Security		ApiKeyAuth
func (s *svcPackageHttpService) handleUpdateServiceTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		svcPackageId := ctx.Param("svcpackage-id")
		if svcPackageId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing svcpackage-id"))
			return
		}
		svcPackageUUID, err := uuid.Parse(svcPackageId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("svcpackage-id invalid (not a uuid)"))
			return
		}

		svcTaskId := ctx.Param("svctask-id")
		if svcTaskId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing stvtask-id"))
			return
		}
		svcTaskUUID, err := uuid.Parse(svcTaskId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("svctask-id invalid (not a uuid)"))
			return
		}

		var dto svcpackagecommands.UpdateServiceTaskDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		dto.SvcTaskId = svcTaskUUID
		dto.SvcPackageId = svcPackageUUID
		if err := s.command.UpdateTask.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

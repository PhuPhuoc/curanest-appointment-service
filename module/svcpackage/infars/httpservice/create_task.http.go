package svcpackagehttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
)

// @Summary		create new service package (staff)
// @Description	create new service package (staff)
// @Tags			service packages
// @Accept			json
// @Produce		json
// @Param			create	form		body					svcpackagecommands.CreateServiceTaskDTO	true	"service task creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/svcpackage/{svcpackage-id}/svctask [post]
// @Security		ApiKeyAuth
func (s *svcPackageHttpService) handleCreateServiceTask() gin.HandlerFunc {
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
		var dto svcpackagecommands.CreateServiceTaskDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreateTask.Handle(ctx.Request.Context(), svcPackageUUID, &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}

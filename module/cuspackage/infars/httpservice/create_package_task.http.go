package cuspackagehttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
)

// @Summary		create customized service package and task (client)
// @Description	create customized service package and task (client)
// @Tags			customized packages
// @Accept			json
// @Produce		json
// @Param			create	form		body					cuspackagecommands.ReqCreatePackageTaskDTO	true	"customized package and task creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/cuspackage [post]
// @Security		ApiKeyAuth
func (s *cusPackageHttpService) handleCreateCustomizedPackageAndTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto cuspackagecommands.ReqCreatePackageTaskDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.cmd.CreateCusPackageAndCusTask.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}

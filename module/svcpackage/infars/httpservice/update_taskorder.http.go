package svcpackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
	"github.com/gin-gonic/gin"
)

//	@Summary		update service task order (staff)
//	@Description	update service task order(staff)
//	@Tags			service packages
//	@Accept			json
//	@Produce		json
//	@Param			create	form		body					svcpackagecommands.UpdateTaskOrderDTO	true	"service task update information"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/api/v1/svcpackage/svctask [patch]
//	@Security		ApiKeyAuth
func (s *svcPackageHttpService) handleUpdateTaskOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto svcpackagecommands.UpdateTaskOrderDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.UpdateTaskOrder.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

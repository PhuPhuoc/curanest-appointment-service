package svcpackagehttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
)

// @Summary		create new service package (staff)
// @Description	create new service package (staff)
// @Tags			service packages
// @Accept			json
// @Produce		json
// @Param			create	form		body					svcpackagecommands.CreateServicePackageDTO	true	"service package creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/svcpackage [post]
// @Security		ApiKeyAuth
func (s *svcPackageHttpService) handleCreateServicePackage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto svcpackagecommands.CreateServicePackageDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreatePackage.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}

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
// @Param			service-id	path		string					true									"service ID (UUID)"
// @Param			create		form		body					svcpackagecommands.ServicePackageDTO	true	"service package creation data"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/services/{service-id}/svcpackage [post]
// @Security		ApiKeyAuth
func (s *svcPackageHttpService) handleCreateServicePackage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		svcId := ctx.Param("service-id")
		if svcId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing service-id"))
			return
		}
		svcUUID, err := uuid.Parse(svcId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a uuid)"))
			return
		}

		var dto svcpackagecommands.ServicePackageDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreatePackage.Handle(ctx.Request.Context(), svcUUID, &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}

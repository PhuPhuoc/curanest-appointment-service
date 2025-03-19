package svcpackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		update service package (staff)
// @Description	update service package (staff)
// @Tags			service packages
// @Accept			json
// @Produce		json
// @Param			service-id		path		string					true										"service ID (UUID)"
// @Param			svcpackage-id	path		string					true										"service-package ID (UUID)"
// @Param			create			form		body					svcpackagecommands.UpdateServicePackageDTO	true	"service package update information"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/services/{service-id}/svcpackage/{svcpackage-id} [put]
// @Security		ApiKeyAuth
func (s *svcPackageHttpService) handleUpdateServicePackage() gin.HandlerFunc {
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

		var dto svcpackagecommands.UpdateServicePackageDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		dto.SvcPackageId = svcpackageUUID
		dto.ServiceId = svcUUID
		if err := s.command.UpdatePackage.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

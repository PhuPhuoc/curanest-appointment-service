package svcpackagehttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

// @Summary		get list service-package by service-id
// @Description	get list service-package by service-id
// @Tags			service packages
// @Accept			json
// @Produce		json
// @Param			service-id	path		string					true	"service ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/services/{service-id}/svcpackage [get]
func (s *svcPackageHttpService) handleGetServicPackage() gin.HandlerFunc {
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

		data, err := s.query.GetServicePackages.Handle(ctx.Request.Context(), svcUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, data)
	}
}

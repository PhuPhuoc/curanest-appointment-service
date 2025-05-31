package svcpackagehttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

//	@Summary		get service-package usage count
//	@Description	get service-package usage count
//	@Tags			service packages
//	@Accept			json
//	@Produce		json
//	@Param			category-id	path		string					true	"service ID (UUID)"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/svcpackage/category/{category-id}/usage-count [get]
func (s *svcPackageHttpService) handleGetServicPackageUsageCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cateId := ctx.Param("category-id")
		if cateId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing category-id"))
			return
		}
		cateUUID, err := uuid.Parse(cateId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("category-id invalid (not a uuid)"))
			return
		}

		data, err := s.query.GetSvcPackageUsageCount.Handle(ctx.Request.Context(), cateUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, data)
	}
}

package invoicehttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

// @Summary		find invoices with cus-package-id
// @Description	find invoices with cus-package-id
// @Tags			invoices
// @Accept			json
// @Produce		json
// @Param			cus-package-id	path		string					true	"customized-package ID (UUID)"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/cuspackage/{cus-package-id}/invoices [get]
// @Security		ApiKeyAuth
func (s *invoiceHttpService) handleFindInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cusPackageId := ctx.Param("cus-package-id")
		if cusPackageId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing cus-package-id"))
			return
		}
		cusPackageUUID, err := uuid.Parse(cusPackageId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("cus-package-id invalid (not a uuid)"))
			return
		}
		invoices, err := s.query.FindInvoice.Handle(ctx.Request.Context(), cusPackageUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, invoices)
	}
}

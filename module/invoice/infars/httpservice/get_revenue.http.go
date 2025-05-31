package invoicehttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/queries"
)

// @Summary		get invoices with patient-ids
// @Description	get invoices with patient-ids
// @Tags			invoices
// @Accept			json
// @Produce		json
// @Param			dates	body		invoicequeries.RequestGetRevenurDTO	true	"List of patient IDs (UUID)"
// @Success		200		{object}	map[string]interface{}				"data"
// @Failure		400		{object}	error								"Bad request error"
// @Router			/api/v1/invoices/revenue [post]
// @Security		ApiKeyAuth
func (s *invoiceHttpService) handleGetRevenue() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req invoicequeries.RequestGetRevenurDTO
		if err := ctx.ShouldBindJSON(&req); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body"))
			return
		}

		invoices, err := s.query.GetRevenue.Handle(ctx.Request.Context(), &req)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, invoices)
	}
}

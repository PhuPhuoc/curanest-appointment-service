package invoicehttpservice

import (
	"strconv"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
)

// @Summary		cannel paytment url
// @Description	cannel paytment url
// @Tags			invoices
// @Accept			json
// @Produce		json
// @Param			order-code	path		string					true	"order code (int)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/invoices/cancel-payment-url/{order-code} [patch]
// @Security		ApiKeyAuth
func (s *invoiceHttpService) handleCancelPaymentUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderCodeStr := ctx.Param("order-code")
		if orderCodeStr == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing order-code"))
			return
		}
		orderCodeInt, err := strconv.ParseInt(orderCodeStr, 0, 64)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("order-code invalid (not a int)"))
			return
		}
		invoice, err := s.query.GetInvoiceByOrderCode.Handle(ctx, orderCodeInt)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		invoiceEntity, _ := invoice.ToInvoiceEntity()
		if err := s.cmd.CancelPaymentUrl.Handle(ctx, invoiceEntity); err != nil {

			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

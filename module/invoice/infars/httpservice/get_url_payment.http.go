package invoicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		get url for payment with payos
//	@Description	get url for payment with payos
//	@Tags			invoices
//	@Accept			json
//	@Produce		json
//	@Param			invoice-id	path		string					true	"invoice ID (UUID)"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/invoices/{invoice-id}/url-payment [get]
//	@Security		ApiKeyAuth
func (s *invoiceHttpService) handleGetUrlPaymentForInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoiceId := ctx.Param("invoice-id")
		if invoiceId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing invoice-id"))
			return
		}
		invoiceUUID, err := uuid.Parse(invoiceId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("nurse-id invalid (not a uuid)"))
			return
		}
		invoice, err := s.query.GetInvoiceById.Handle(ctx.Request.Context(), invoiceUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		invoiceDTO := &invoicecommands.DetailInvoiceCmdDTO{
			Id:       invoice.Id,
			TotalFee: invoice.TotalFee,
		}
		url, err := s.cmd.GetUrlPayment.Handle(ctx.Request.Context(), invoiceDTO)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, url)
	}
}

package invoicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
			Id:        invoice.Id,
			OrderCode: invoice.OrderCode,
			TotalFee:  invoice.TotalFee,
		}
		url, err := s.cmd.GetUrlPayment.Handle(ctx.Request.Context(), invoiceDTO)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, url)
	}
}

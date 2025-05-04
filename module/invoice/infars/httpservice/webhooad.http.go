package invoicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/commands"
	"github.com/gin-gonic/gin"
)

// @Summary		Handle PayOS webhook
// @Description	Process webhook notifications from PayOS to update invoice payment status
// @Tags			invoices
// @Accept			json
// @Produce		json
// @Param			webhook_data	body		invoicecommands.PayosWebhookData	true	"PayOS webhook data"
// @Success		200				{object}	map[string]interface{}				"Success response"
// @Failure		400				{object}	map[string]interface{}				"Bad request error"
// @Router			/api/v1/invoices/webhooks [post]
func (s *invoiceHttpService) handlePayosWebhook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto invoicecommands.PayosWebhookData

		// Bind JSON body
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("Invalid request payload"))
			return
		}

		invoiceDTO, _ := s.query.GetInvoiceByOrderCode.Handle(ctx.Request.Context(), dto.Data.OrderCode)
		invoiceEntity, _ := invoiceDTO.ToInvoiceEntity()

		// Call webhook handler
		err := s.cmd.WebHookGoong.Handle(ctx.Request.Context(), s.checksumKey, &dto, invoiceEntity)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		// Respond to PayOS
		common.ResponseUpdated(ctx)
	}
}

package invoicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		create new paytment url
// @Description	create new paytment url
// @Tags			invoices
// @Accept			json
// @Produce		json
// @Param			invoice-id	path		string					true	"invoice-id uuid"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/invoices/{invoice-id}/create-payment-url [patch]
// @Security		ApiKeyAuth
func (s *invoiceHttpService) handleCreateNewPaymentUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoiceId := ctx.Param("invoice-id")
		if invoiceId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing invoice-id"))
			return
		}
		invoiceUUID, err := uuid.Parse(invoiceId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invoice-id invalid (not a uuid)"))
			return
		}
		invoice, err := s.query.GetInvoiceById.Handle(ctx, invoiceUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		invoiceEntity, _ := invoice.ToInvoiceEntity()
		if err := s.cmd.CreateNewUrl.Handle(ctx, invoiceEntity); err != nil {

			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}

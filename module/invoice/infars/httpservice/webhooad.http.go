package invoicehttpservice

import (
	"net/http"

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
	return func(c *gin.Context) {
		var dto invoicecommands.PayosWebhookData

		// Bind JSON body
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Call webhook handler
		err := s.cmd.WebHookGoong.Handle(c.Request.Context(), s.checksumKey, &dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Respond to PayOS
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

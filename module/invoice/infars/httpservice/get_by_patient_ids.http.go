package invoicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		get invoices with patient-ids
// @Description	get invoices with patient-ids
// @Tags			invoices
// @Accept			json
// @Produce		json
// @Param			patientIds	body		invoicequeries.RequestGetInvoicesByPatientIds	true	"List of patient IDs (UUID)"
// @Success		200			{object}	map[string]interface{}							"data"
// @Failure		400			{object}	error											"Bad request error"
// @Router			/api/v1/invoices/patient [post]
// @Security		ApiKeyAuth
func (s *invoiceHttpService) handleInvoicesByPatientIds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req invoicequeries.RequestGetInvoicesByPatientIds
		if err := ctx.ShouldBindJSON(&req); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body"))
			return
		}

		if len(req.PatientIds) == 0 {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("patientIds is empty"))
			return
		}

		invoices, err := s.query.GetInvoiceByPatientIds.Handle(ctx.Request.Context(), req.PatientIds)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, invoices)
	}
}

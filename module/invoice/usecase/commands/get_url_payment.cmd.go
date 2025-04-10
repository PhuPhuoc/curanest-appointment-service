package invoicecommands

import (
	"context"

	"github.com/payOSHQ/payos-lib-golang"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getUrlPaymentHandler struct {
	cmdRepo     InvoiceCommandRepo
	payosConfig common.PayOSConfig
}

func NewGetUrlPaymentHandler(cmdRepo InvoiceCommandRepo, payosConfig common.PayOSConfig) *getUrlPaymentHandler {
	return &getUrlPaymentHandler{
		cmdRepo:     cmdRepo,
		payosConfig: payosConfig,
	}
}

func (h *getUrlPaymentHandler) Handle(ctx context.Context, invoice *DetailInvoiceCmdDTO) (*UrlPayment, error) {
	payos.Key(h.payosConfig.ClientId, h.payosConfig.ApiKey, h.payosConfig.CheckSumKey)

	paymentRequest := payos.CheckoutRequestType{
		OrderCode:   invoice.OrderCode,
		Amount:      int(invoice.TotalFee),
		Description: "curanest",
		CancelUrl:   "http://localhost:8080/cancel",
		ReturnUrl:   "http://localhost:8080/success",
	}

	response, err := payos.CreatePaymentLink(paymentRequest)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("failed to create payment url").
			WithInner(err.Error())
	}

	url := &UrlPayment{
		Url: response.CheckoutUrl,
	}

	return url, nil
}

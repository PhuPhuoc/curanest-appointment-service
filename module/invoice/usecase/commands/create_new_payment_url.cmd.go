package invoicecommands

import (
	"context"
	"math/rand"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	"github.com/payOSHQ/payos-lib-golang"
)

type createNewPaymentUrlHandler struct {
	cmdRepo     InvoiceCommandRepo
	payosConfig common.PayOSConfig
}

func NewCreateNewPaymentUrlHandler(cmdRepo InvoiceCommandRepo, payosConfig common.PayOSConfig) *createNewPaymentUrlHandler {
	return &createNewPaymentUrlHandler{
		cmdRepo:     cmdRepo,
		payosConfig: payosConfig,
	}
}

func (h *createNewPaymentUrlHandler) Handle(ctx context.Context, e *invoicedomain.Invoice) error {
	orderCode := time.Now().Unix()*1000 + int64(rand.Intn(1000))
	payos.Key(h.payosConfig.ClientId, h.payosConfig.ApiKey, h.payosConfig.CheckSumKey)

	paymentRequest := payos.CheckoutRequestType{
		OrderCode:   orderCode,
		Amount:      int(e.GetTotalFee()),
		Description: "curanest",
		CancelUrl:   "https://curanest.com.vn/payment-result-fail",
		ReturnUrl:   "https://curanest.com.vn/payment-result-success",
	}

	response, err := payos.CreatePaymentLink(paymentRequest)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("failed to create payment url").
			WithInner(err.Error())
	}

	newInvoice, _ := invoicedomain.NewInvoice(
		e.GetID(),
		e.GetCusPackageID(),
		&orderCode,
		e.GetTotalFee(),
		e.GetPaymentStatus(),
		e.GetNote(),
		&response.CheckoutUrl,
		&response.QRCode,
		e.GetCreatedAt(),
	)

	if err := h.cmdRepo.UpdateInvoice(ctx, newInvoice); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update this invoice").
			WithInner(err.Error())
	}

	return nil
}

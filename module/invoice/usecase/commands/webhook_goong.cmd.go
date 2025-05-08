package invoicecommands

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type webhookGoongHandler struct {
	cmdRepo           InvoiceCommandRepo
	cuspackageFetcher CusPackageFetcher
}

func NewWebhoobGoongHandler(cmdRepo InvoiceCommandRepo, cuspackageFetcher CusPackageFetcher) *webhookGoongHandler {
	return &webhookGoongHandler{
		cmdRepo:           cmdRepo,
		cuspackageFetcher: cuspackageFetcher,
	}
}

func (h *webhookGoongHandler) Handle(ctx context.Context, checkSumKey string, dto *PayosWebhookData, invoiceEntity *invoicedomain.Invoice) error {
	log.Println("dto gooong: ", dto)
	// Check if transaction is successful
	if !dto.Success {
		log.Printf("Webhook not successful: Success=%v", dto.Success)
		return nil
	}

	// Check if the transaction is successful based on data fields
	if dto.Data.Code != "00" || dto.Data.Desc != "success" || dto.Data.AccountNumber == "" || dto.Data.TransactionDateTime == "" {
		log.Printf("Transaction not paid: Code=%v, Desc=%v, AccountNumber=%v, TransactionDateTime=%v",
			dto.Data.Code, dto.Data.Desc, dto.Data.AccountNumber, dto.Data.TransactionDateTime)
		return nil
	}

	// Call repository to update invoice
	orderCode := fmt.Sprintf("%d", dto.Data.OrderCode)
	log.Printf("Updating invoice with orderCode: %s", orderCode)
	if err := h.cmdRepo.UpdateInvoiceFromGoong(ctx, orderCode); err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}

	cusPackage, err := h.cuspackageFetcher.FindCusPackage(ctx, invoiceEntity.GetCusPackageID())
	if err == nil {
		cusPaymentStatus := cuspackagedomain.PaymentStatusUnpaid
		totalUnpaidAmount := cusPackage.GetTotalFee() - invoiceEntity.GetTotalFee()
		if totalUnpaidAmount <= 0 {
			cusPaymentStatus = cuspackagedomain.PaymentStatusPaid
		}

		updateCusPackage, _ := cuspackagedomain.NewCustomizedPackage(
			cusPackage.GetID(),
			cusPackage.GetServicePackageID(),
			cusPackage.GetPatientID(),
			cusPackage.GetName(),
			cusPackage.GetTotalFee(),
			invoiceEntity.GetTotalFee(),
			totalUnpaidAmount,
			cusPaymentStatus,
			cusPackage.GetCreatedAt(),
		)

		_ = h.cuspackageFetcher.UpdateCustomizedPackage(ctx, updateCusPackage)

	}

	log.Printf("Invoice updated successfully for orderCode: %s", orderCode)
	return nil
}

// verifySignature verifies the PayOS webhook signature
func (h *webhookGoongHandler) verifySignature(checksumKey string, dto *PayosWebhookData) bool {
	receivedSignature := dto.Signature
	dto.Signature = "" // Remove signature for verification

	// Convert data to JSON
	dataBytes, err := json.Marshal(dto)
	if err != nil {
		return false
	}

	// Sort keys for consistent signature
	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataBytes, &dataMap); err != nil {
		return false
	}
	keys := make([]string, 0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedData := make(map[string]interface{})
	for _, k := range keys {
		sortedData[k] = dataMap[k]
	}
	sortedBytes, err := json.Marshal(sortedData)
	if err != nil {
		return false
	}

	// Calculate HMAC SHA256
	mac := hmac.New(sha256.New, []byte(checksumKey))
	mac.Write(sortedBytes)
	calculatedSignature := fmt.Sprintf("%x", mac.Sum(nil))

	return calculatedSignature == receivedSignature
}

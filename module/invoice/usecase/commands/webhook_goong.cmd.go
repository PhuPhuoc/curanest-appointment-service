package invoicecommands

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
)

type webhookGoongHandler struct {
	cmdRepo InvoiceCommandRepo
}

func NewWebhoobGoongHandler(cmdRepo InvoiceCommandRepo) *webhookGoongHandler {
	return &webhookGoongHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *webhookGoongHandler) Handle(ctx context.Context, checkSumKey string, dto *PayosWebhookData) error {
	// Verify signature
	// if !h.verifySignature(checkSumKey, dto) {
	// 	return fmt.Errorf("invalid signature")
	// }

	// Check if transaction is successful
	if !dto.Success || dto.Data.Status != "PAID" {
		return nil // Not a paid transaction, no error but no update
	}

	// Call repository to update invoice
	orderCode := fmt.Sprintf("%d", dto.Data.OrderCode)
	if err := h.cmdRepo.UpdateInvoiceFromGoong(ctx, orderCode); err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}

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

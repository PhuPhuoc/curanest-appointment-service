package invoicerepository

import (
	"context"
	"fmt"
)

func (repo *invoiceRepo) UpdateInvoiceFromGoong(ctx context.Context, orderCode string) error {
	query := `
        UPDATE invoices 
        SET payment_status = 'paid', note = 'da thanh toan xong' 
        WHERE order_code = ? AND payment_status = 'unpaid'
    `

	// Execute query with context
	result, err := repo.db.ExecContext(ctx, query, orderCode)
	if err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}

	// Check affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no invoice found with order_code %s or invoice already paid", orderCode)
	}

	return nil
}

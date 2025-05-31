package invoicerepository

import "context"

func (repo *invoiceRepo) GetTotalRevenueInDuration(ctx context.Context, dateFrom, dateTo string) (float64, error) {
	query := `
		SELECT COALESCE(SUM(total_fee), 0)
		FROM invoices
		WHERE payment_status = 'paid'
		  AND DATE(created_at) BETWEEN ? AND ?
	`

	var totalRevenue float64
	err := repo.db.GetContext(ctx, &totalRevenue, query, dateFrom, dateTo)
	if err != nil {
		return 0, err
	}

	return totalRevenue, nil
}

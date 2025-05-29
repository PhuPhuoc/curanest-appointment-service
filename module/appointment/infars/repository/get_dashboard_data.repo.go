package appointmentrepository

import (
	"context"
	"fmt"
	"strings"

	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

func (repo *appointmentRepo) GetDashboardData(ctx context.Context, filter *appointmentqueries.FilterDashboardDTO) (int, int, int, error) {
	var waitingCount, upcomingCount, dateRangeCount int

	fmt.Println("filter: ", filter)

	// Base query parts
	baseQuery := `
		SELECT COUNT(*) FROM appointments a
		JOIN services s ON a.service_id = s.id
		WHERE a.deleted_at IS NULL AND s.deleted_at IS NULL
	`
	var args []interface{}
	var conditions []string

	// Add category filter if not admin
	if !filter.IsAdmin && filter.CategoryId != nil {
		conditions = append(conditions, "s.category_id = ?")
		args = append(args, filter.CategoryId.String())
	}

	// Query 1: Waiting appointments
	queryWaiting := baseQuery + " AND a.status = 'waiting'"
	if len(conditions) > 0 {
		queryWaiting += " AND " + strings.Join(conditions, " AND ")
	}
	if err := repo.db.GetContext(ctx, &waitingCount, queryWaiting, args...); err != nil {
		return 0, 0, 0, err
	}

	// Query 2: Upcoming appointments
	queryUpcoming := baseQuery + " AND a.status = 'upcoming'"
	if len(conditions) > 0 {
		queryUpcoming += " AND " + strings.Join(conditions, " AND ")
	}
	if err := repo.db.GetContext(ctx, &upcomingCount, queryUpcoming, args...); err != nil {
		return 0, 0, 0, err
	}

	queryDateRange := baseQuery + " AND a.est_date BETWEEN ? AND ?"
	if len(conditions) > 0 {
		queryDateRange += " AND " + strings.Join(conditions, " AND ")
	}

	argsWithDates := []interface{}{filter.DateFrom, filter.DateTo}
	if len(args) > 0 {
		argsWithDates = append(argsWithDates, args...)
	}

	if err := repo.db.GetContext(ctx, &dateRangeCount, queryDateRange, argsWithDates...); err != nil {
		return 0, 0, 0, err
	}

	return waitingCount, upcomingCount, dateRangeCount, nil
}

package appointmentrepository

import (
	"context"

	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

func (repo *appointmentRepo) GetDashboardData(ctx context.Context, filter *appointmentqueries.FilterDashboardDTO) (int, int, int, error) {
	return 10, 10, 10, nil
}

package appointmentqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getItemDashboardHandler struct {
	queryRepo AppointmentQueryRepo
}

func NewGetItemDashboardHandler(queryRepo AppointmentQueryRepo) *getItemDashboardHandler {
	return &getItemDashboardHandler{
		queryRepo: queryRepo,
	}
}

func (h *getItemDashboardHandler) Handle(ctx context.Context, filter *FilterDashboardDTO) (*ItemDashboardDTO, error) {
	upComing, waiting, total, err := h.queryRepo.GetDashboardData(ctx, filter)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get dashboard data").
			WithInner(err.Error())
	}

	dashboard := &ItemDashboardDTO{
		TotalService: 100,
		UpcomingApps: upComing,
		WaitingApps:  waiting,
		TotalApps:    total,
	}

	return dashboard, nil
}

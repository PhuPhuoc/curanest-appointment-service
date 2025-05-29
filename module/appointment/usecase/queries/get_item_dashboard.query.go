package appointmentqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getItemDashboardHandler struct {
	queryRepo      AppointmentQueryRepo
	serviceFetcher ServiceFetcher
}

func NewGetItemDashboardHandler(queryRepo AppointmentQueryRepo, serviceFetcher ServiceFetcher) *getItemDashboardHandler {
	return &getItemDashboardHandler{
		queryRepo:      queryRepo,
		serviceFetcher: serviceFetcher,
	}
}

func (h *getItemDashboardHandler) Handle(ctx context.Context, filter *FilterDashboardDTO) (*ItemDashboardDTO, error) {
	waiting, upComing, total, err := h.queryRepo.GetDashboardData(ctx, filter)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get dashboard data").
			WithInner(err.Error())
	}

	totalService, err := h.serviceFetcher.GetCountTotalService(ctx, filter.CategoryId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get dashboard data - total service").
			WithInner(err.Error())
	}
	if totalService == 0 {
		totalService = 20
	}

	dashboard := &ItemDashboardDTO{
		TotalService: totalService,
		UpcomingApps: upComing,
		WaitingApps:  waiting,
		TotalApps:    total,
	}

	return dashboard, nil
}

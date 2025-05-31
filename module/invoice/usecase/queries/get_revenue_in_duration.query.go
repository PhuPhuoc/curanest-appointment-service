package invoicequeries

import (
	"context"
	"fmt"
	"strings"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getRevenueHandler struct {
	queryRepo InvoiceQueryRepo
}

func NewGetRevenueHandler(queryRepo InvoiceQueryRepo) *getRevenueHandler {
	return &getRevenueHandler{
		queryRepo: queryRepo,
	}
}

func (h *getRevenueHandler) Handle(ctx context.Context, dates *RequestGetRevenurDTO) ([]RevenurDTO, error) {
	var result []RevenurDTO

	for _, rangeStr := range dates.Dates {
		parts := strings.Split(rangeStr, "/")
		if len(parts) != 2 {
			err_mess := fmt.Errorf("invalid date range format: %s", rangeStr)
			return []RevenurDTO{}, common.NewInternalServerError().
				WithReason("cannot get revenue").
				WithInner(err_mess.Error())
		}

		dateFrom := strings.TrimSpace(parts[0])
		dateTo := strings.TrimSpace(parts[1])

		revenue, err := h.queryRepo.GetTotalRevenueInDuration(ctx, dateFrom, dateTo)
		if err != nil {
			return []RevenurDTO{}, common.NewInternalServerError().
				WithReason("cannot get revenue").
				WithInner(err.Error())
		}

		result = append(result, RevenurDTO{
			Date:    rangeStr,
			Revenue: revenue,
		})
	}

	return result, nil
}

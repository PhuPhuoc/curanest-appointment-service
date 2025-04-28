package appointmentqueries

import (
	"context"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type checkNursesAvailabilityHandler struct {
	queryRepo AppointmentQueryRepo
}

func NewCheckNursesAvailabilityHandler(queryRepo AppointmentQueryRepo) *checkNursesAvailabilityHandler {
	return &checkNursesAvailabilityHandler{
		queryRepo: queryRepo,
	}
}

func (h *checkNursesAvailabilityHandler) Handle(ctx context.Context, dto *CheckNursesAvailabilityRequestDTO) ([]NurseDateMappingResult, error) {
	response := make([]NurseDateMappingResult, len(dto.NursesDates))
	for i, obj := range dto.NursesDates {
		respObj := NurseDateMappingResult{
			NurseId:        obj.NurseId,
			EstStartDate:   obj.EstStartDate,
			EstDuration:    obj.EstDuration,
			IsAvailability: true,
		}

		beginDate := time.Date(obj.EstStartDate.Year(), obj.EstStartDate.Month(), obj.EstStartDate.Day(), 0, 0, 0, 0, obj.EstStartDate.Location())
		endDate := time.Date(beginDate.Year(), beginDate.Month(), beginDate.Day(), 23, 59, 0, 0, obj.EstStartDate.Location())
		endDate = endDate.Add(time.Duration(obj.EstDuration+20) * time.Minute)

		entities, err := h.queryRepo.GetAppointmentInADayOfNursing(ctx, obj.NurseId, beginDate, endDate)
		if err != nil {
			return []NurseDateMappingResult{}, common.NewInternalServerError().
				WithReason("cannot get appointment in a day of this nursing").
				WithInner(err.Error())
		}

		if len(entities) != 0 {
			estTravelTime := 20
			currentStartDate := obj.EstStartDate
			currentEndDate := currentStartDate.Add(time.Duration(obj.EstDuration+estTravelTime) * time.Minute)

			for _, app := range entities {
				appStartDate := app.GetEstDate()
				appEndDate := appStartDate.Add(time.Duration(app.GetTotalEstDuration()+estTravelTime) * time.Minute)

				if isOverlapping(appStartDate, appEndDate, currentStartDate, currentEndDate) {
					respObj.IsAvailability = false
				}
			}
		}
		response[i] = respObj
	}

	return response, nil
}

func isOverlapping(start1, end1, start2, end2 time.Time) bool {
	return !(end1.Before(start2) || end2.Before(start1))
}

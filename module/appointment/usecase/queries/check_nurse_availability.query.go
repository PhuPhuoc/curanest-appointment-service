package appointmentqueries

import (
	"context"
	"fmt"
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

		estEndDate := obj.EstStartDate.Add(time.Duration(obj.EstDuration+20) * time.Minute)
		if err := h.queryRepo.IsNurseAvailability(ctx, obj.NurseId, obj.EstStartDate, estEndDate); err != nil {
			if err == common.ErrNurseNotAvailable {
				respObj.IsAvailability = false
			} else {
				return nil, common.NewInternalServerError().
					WithReason(fmt.Sprintf("cannot check availability for nurse with id: %v", obj.NurseId.String())).
					WithInner(err.Error())
			}
		}
		response[i] = respObj
	}

	return response, nil
}

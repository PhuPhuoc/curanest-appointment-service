package appointmentqueries

import (
	"context"
	"fmt"

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
			Date:           obj.Date,
			IsAvailability: true,
		}
		fmt.Println("respObj: ", respObj)

		if err := h.queryRepo.IsNurseAvailability(ctx, obj.NurseId, obj.Date); err != nil {
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

	fmt.Println("resp: ", response)

	return response, nil
}

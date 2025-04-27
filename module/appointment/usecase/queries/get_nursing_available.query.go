package appointmentqueries

import (
	"context"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type getNursingAvailableHandler struct {
	queryRepo      AppointmentQueryRepo
	nurseServceApi NursingServiceExternalAPI
}

func NewGetNursingAvailableHandler(queryRepo AppointmentQueryRepo, nurseServceApi NursingServiceExternalAPI) *getNursingAvailableHandler {
	return &getNursingAvailableHandler{
		queryRepo:      queryRepo,
		nurseServceApi: nurseServceApi,
	}
}

func (h *getNursingAvailableHandler) Handle(ctx context.Context, serviceId uuid.UUID, estStartDate time.Time, estDuration int) ([]NurseDTO, error) {
	estEndDate := estStartDate.Add(time.Duration(estDuration+20) * time.Minute)

	apps, err := h.queryRepo.GetAppointmentInDate(ctx, estStartDate, estEndDate)
	if err != nil {
		return []NurseDTO{}, common.NewInternalServerError().
			WithReason("cannot get list appointments").
			WithInner(err.Error())
	}

	nurses, err := h.nurseServceApi.GetNursingByServiceIdRPC(ctx, serviceId)
	if err != nil {
		return []NurseDTO{}, common.NewInternalServerError().
			WithReason("cannot get list nurses").
			WithInner(err.Error())
	}

	availableNurses := []NurseDTO{}
	for _, nurse := range nurses {
		flagAvailable := true
		for _, app := range apps {
			if nurse.NurseId == *app.GetNursingID() {
				flagAvailable = false
				continue
			}
		}
		if flagAvailable {
			availableNurses = append(availableNurses, nurse)
		}
	}

	return availableNurses, nil
}

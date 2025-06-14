package appointmentqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type getNursingTimeSheetHandler struct {
	queryRepo AppointmentQueryRepo
}

func NewGetNursingTimeSheetHandler(queryRepo AppointmentQueryRepo) *getNursingTimeSheetHandler {
	return &getNursingTimeSheetHandler{
		queryRepo: queryRepo,
	}
}

func (h *getNursingTimeSheetHandler) Handle(ctx context.Context, filter *FilterGetNursingTimesheetDTO) ([]TimesheetDTO, error) {
	flagPaing := false
	filterQuery := FilterGetAppointmentDTO{
		NursingId:   &filter.NursingId,
		EstDateFrom: filter.EstDateFrom,
		EstDateTo:   filter.EstDateTo,
		ApplyPaging: &flagPaing,
	}
	entities, err := h.queryRepo.GetAppointment(ctx, &filterQuery)
	if err != nil {
		return []TimesheetDTO{}, common.NewInternalServerError().
			WithReason("cannot get appointment").
			WithInner(err.Error())
	}
	if len(entities) == 0 {
		return []TimesheetDTO{}, nil
	}

	estTravelTime := 20 // default 20 min - change later
	dtos := make([]TimesheetDTO, len(entities))
	for i, entity := range entities {
		if entity.GetStatus() != appointmentdomain.AppStatusCancel {
			dtos[i] = *toTimesheetDTO(&entity, estTravelTime)
		}
	}

	return dtos, nil
}

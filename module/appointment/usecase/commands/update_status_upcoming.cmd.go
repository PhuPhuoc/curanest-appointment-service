package apppointmentcommands

import (
	"context"
	"fmt"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type updateStatusUpcomingHandler struct {
	cmdRepo  AppointmentCommandRepo
	goongApi ExternalGoongAPI
}

func NewUpdateStatusUpcomingHandler(cmdRepo AppointmentCommandRepo, goongApi ExternalGoongAPI) *updateStatusUpcomingHandler {
	return &updateStatusUpcomingHandler{
		cmdRepo:  cmdRepo,
		goongApi: goongApi,
	}
}

func (h *updateStatusUpcomingHandler) Handle(ctx context.Context, originCode string, entity *appointmentdomain.Appointment) error {
	var distStr string
	var duraVal int
	if entity.GetNursingID() == nil {
		return common.NewInternalServerError().
			WithReason("cannot update appointment's status to upcoming. This appointment don't have nursing")
	}
	goongApiResp, err := h.goongApi.GetDistanceFromGoong(ctx, originCode, entity.GetPatientLatLng())
	if err != nil {
		distStr = "cannot find location"
		duraVal = 1200
	} else {
		distStr = goongApiResp.Rows[0].Elements[0].Distance.Text
		duraVal = goongApiResp.Rows[0].Elements[0].Duration.Value
	}

	actDate := entity.GetEstDate().Add(time.Duration(duraVal) * time.Second)

	hours := duraVal / 3600
	minutes := (duraVal % 3600) / 60

	var etaText string
	if hours > 0 && minutes > 0 {
		etaText = fmt.Sprintf("%d hours %d minutes", hours, minutes)
	} else if hours > 0 {
		etaText = fmt.Sprintf("%d hours", hours)
	} else {
		etaText = fmt.Sprintf("%d minutes", minutes)
	}

	fmt.Printf(
		"The nurse is currently %v kilometers away from you and is expected to arrive in about %v.\n"+
			"The service appointment is scheduled to take place at %02d:%02d on %s %d, %d.\n",
		distStr,
		etaText,
		actDate.Hour(),
		actDate.Minute(),
		actDate.Month().String(),
		actDate.Day(),
		actDate.Year(),
	)

	updateEntity, _ := appointmentdomain.NewAppointment(
		entity.GetID(),
		entity.GetServiceID(),
		entity.GetCusPackageID(),
		entity.GetPatientID(),
		entity.GetNursingID(),
		entity.GetPatientAddress(),
		entity.GetPatientLatLng(),
		appointmentdomain.AppStatusUpcoming,
		entity.GetTotalEstDuration(),
		entity.GetEstDate(),
		&actDate,
		entity.GetCreatedAt(),
	)
	if err := h.cmdRepo.UpdateAppointment(ctx, updateEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update appointment's status").
			WithInner(err.Error())
	}
	return nil
}

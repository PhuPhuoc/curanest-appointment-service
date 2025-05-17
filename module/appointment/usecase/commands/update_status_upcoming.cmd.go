package apppointmentcommands

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type updateStatusUpcomingHandler struct {
	cmdRepo         AppointmentCommandRepo
	pushNotiFetcher ExternalPushNotiService
}

func NewUpdateStatusUpcomingHandler(cmdRepo AppointmentCommandRepo, pushNotiFetcher ExternalPushNotiService) *updateStatusUpcomingHandler {
	return &updateStatusUpcomingHandler{
		cmdRepo:         cmdRepo,
		pushNotiFetcher: pushNotiFetcher,
	}
}

func (h *updateStatusUpcomingHandler) Handle(ctx context.Context, entity *appointmentdomain.Appointment) error {
	if entity.GetNursingID() == nil {
		return common.NewInternalServerError().
			WithReason("cannot update appointment's status to upcoming. This appointment doesn't have nursing")
	}

	now := time.Now()

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
		&now,
		entity.GetCreatedAt(),
	)

	if err := h.cmdRepo.UpdateAppointment(ctx, updateEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update appointment's status").
			WithInner(err.Error())
	}

	nursingName := "A hihi"
	patientame := "B huhu"

	diff := entity.GetEstDate().Sub(now)
	minutesUntil := int(diff.Minutes())
	contentVi := fmt.Sprintf(
		"Hệ thống thông báo: \n"+
			"Điều dưỡng %v đang trên đường đến cuộc hẹn với bệnh nhân %s.\n"+
			"Cuộc hẹn dự kiến bắt đầu sau %s phút nữa.\n",
		nursingName,
		patientame,
		minutesUntil,
	)

	// Gửi thông báo
	reqPushNoti := common.PushNotiRequest{
		AccountID: *entity.GetNursingID(),
		Content:   contentVi,
		Route:     "/(tabs)/schedule",
	}
	if err_noti := h.pushNotiFetcher.PushNotification(ctx, &reqPushNoti); err_noti != nil {
		log.Println("error push noti for nursing: ", err_noti)
	}

	return nil
}

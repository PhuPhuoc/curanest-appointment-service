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
	goongApi        ExternalGoongAPI
	pushNotiFetcher ExternalPushNotiService
}

func NewUpdateStatusUpcomingHandler(cmdRepo AppointmentCommandRepo, goongApi ExternalGoongAPI, pushNotiFetcher ExternalPushNotiService) *updateStatusUpcomingHandler {
	return &updateStatusUpcomingHandler{
		cmdRepo:         cmdRepo,
		goongApi:        goongApi,
		pushNotiFetcher: pushNotiFetcher,
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

	now := time.Now()
	actDate := now.Add(time.Duration(duraVal) * time.Second)

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

	updateEntity, _ := appointmentdomain.NewAppointment(
		entity.GetID(),
		entity.GetServiceID(),
		entity.GetSvcpackageID(),
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

	/*
		content := fmt.Sprintf(
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
	*/

	actDateVN := actDate
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err == nil {
		actDateVN = actDate.In(loc)
	}

	contentVi := fmt.Sprintf(
		"Y tá hiện đang cách bạn khoảng %v km và dự kiến sẽ đến trong khoảng %v.\n"+
			"Cuộc hẹn dịch vụ được lên lịch vào lúc %s.\n",
		distStr,
		etaText,
		actDateVN.Format("15:04 ngày 02 tháng 01 năm 2006"),
	)
	reqPushNoti := common.PushNotiRequest{
		AccountID: *entity.GetNursingID(),
		Content:   contentVi,
		Route:     "/(tabs)/schedule",
	}
	err_noti := h.pushNotiFetcher.PushNotification(ctx, &reqPushNoti)
	if err_noti != nil {
		log.Println("error push noti for nursing: ", err_noti)
	}
	return nil
}

package apppointmentcommands

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type updateAppointmentHandler struct {
	cmdRepo      AppointmentCommandRepo
	cusTaskFetch CusTaskFetcher
	recordFetch  MedicalRecordFetcher
}

func NewUpdateAppointmentStatusHandler(cmdRepo AppointmentCommandRepo, cusTaskFetch CusTaskFetcher, recordFetch MedicalRecordFetcher) *updateAppointmentHandler {
	return &updateAppointmentHandler{
		cmdRepo:      cmdRepo,
		cusTaskFetch: cusTaskFetch,
		recordFetch:  recordFetch,
	}
}

func (h *updateAppointmentHandler) Handle(ctx context.Context, newStatus appointmentdomain.AppointmentStatus, entity *appointmentdomain.Appointment) error {
	currentStatus := entity.GetStatus()
	if currentStatus == newStatus {
		return common.NewBadRequestError().
			WithReason(fmt.Sprintf("The current status is already: %v", currentStatus))
	}
	if currentStatus == appointmentdomain.AppStatusSuccess {
		return common.NewBadRequestError().
			WithReason("the status of the appointment is already 'success' and cannot be changed")
	}
	if currentStatus == appointmentdomain.AppStatusUpcoming && newStatus != appointmentdomain.AppStatusSuccess {
		return common.NewBadRequestError().
			WithReason("the current status of the appointment is 'upcoming' and can only be confirmed as 'success'")
	}

	switch newStatus {
	case appointmentdomain.AppStatusSuccess:
		if err := h.isTransitionToSuccessValid(ctx, entity); err != nil {
			return err
		}
	case appointmentdomain.AppStatusUpcoming:
		if err := h.isTransitionToUpcomingValid(ctx, entity); err != nil {
			return err
		}
	default:
		return common.NewBadRequestError().
			WithReason(fmt.Sprintf("the status you are trying to update (%v) this appointment to is invalid", newStatus))
	}

	updateEntity, _ := appointmentdomain.NewAppointment(
		entity.GetID(),
		entity.GetServiceID(),
		entity.GetCusPackageID(),
		entity.GetPatientID(),
		entity.GetNursingID(),
		entity.GetPatientAddress(),
		entity.GetPatientLatLng(),
		newStatus,
		entity.GetTotalEstDuration(),
		entity.GetEstDate(),
		entity.GetActDate(),
		entity.GetCreatedAt(),
	)

	if err := h.cmdRepo.UpdateAppointment(ctx, updateEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update appointment's status").
			WithInner(err.Error())
	}
	return nil
}

func (h *updateAppointmentHandler) isTransitionToSuccessValid(ctx context.Context, entity *appointmentdomain.Appointment) error {
	if entity.GetStatus() != appointmentdomain.AppStatusUpcoming {
		return common.NewBadRequestError().
			WithReason("you can only change the status to 'success' when the current status is 'upcoming'")
	} else {
		if err := h.cusTaskFetch.CheckCusTasksAllDone(ctx, entity.GetCusPackageID()); err != nil {
			return common.NewBadRequestError().
				WithReason("the tasks have not been fully completed, so the appointment status cannot be changed to 'success'")
		}
		if err := h.recordFetch.CheckMedicalRecordDone(ctx, entity.GetCusPackageID()); err != nil {
			return common.NewBadRequestError().
				WithReason("the medical record has not been completed yet, so the appointment status cannot be changed to 'success'")
		}
	}
	return nil
}

func (h *updateAppointmentHandler) isTransitionToUpcomingValid(ctx context.Context, entity *appointmentdomain.Appointment) error {
	if entity.GetNursingID() == nil {
		return common.NewBadRequestError().
			WithReason("this appointment does not have an assigned nurse yet, so the status cannot be updated to 'upcoming'")
	}
	return nil
}

func (h *updateAppointmentHandler) isTransitionToChangedValid(ctx context.Context, entity *appointmentdomain.Appointment) error {
	if entity.GetNursingID() == nil {
		return common.NewBadRequestError().
			WithReason("this appointment does not have an assigned nurse yet, so the status cannot be updated to 'changed'")
	}
	return nil
}

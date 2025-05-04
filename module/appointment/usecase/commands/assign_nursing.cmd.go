package apppointmentcommands

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type assignNursingHandler struct {
	cmdRepo              AppointmentCommandRepo
	medicalRecordFetcher MedicalRecordFetcher
	txManager            common.TransactionManager
}

func NewAssignNursingHandler(cmdRepo AppointmentCommandRepo, txManager common.TransactionManager, medicalRecordFetcher MedicalRecordFetcher) *assignNursingHandler {
	return &assignNursingHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *assignNursingHandler) Handle(ctx context.Context, nursingId uuid.UUID, entity *appointmentdomain.Appointment) error {
	ctx, err := h.txManager.Begin(ctx)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot start transaction").
			WithInner(err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			h.txManager.Rollback(ctx)
			panic(p)
		} else if err != nil {
			h.txManager.Rollback(ctx)
		}
	}()

	updateEntity, _ := appointmentdomain.NewAppointment(
		entity.GetID(),
		entity.GetServiceID(),
		entity.GetSvcpackageID(),
		entity.GetCusPackageID(),
		entity.GetPatientID(),
		&nursingId,
		entity.GetPatientAddress(),
		entity.GetPatientLatLng(),
		appointmentdomain.AppStatusConfirmed,
		entity.GetTotalEstDuration(),
		entity.GetEstDate(),
		entity.GetActDate(),
		entity.GetCreatedAt(),
	)
	if err = h.cmdRepo.UpdateAppointment(ctx, updateEntity); err != nil {
		return common.NewBadRequestError().
			WithReason(fmt.Sprintf("cannot assign nursing to appointment: %v", err)).
			WithInner(err.Error())
	}

	medicalReEntity, err := h.medicalRecordFetcher.FindMedicalRecordByAppsId(ctx, entity.GetID())
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot found medical record of this appointment").
			WithInner(err.Error())
	}

	updateMR, _ := cuspackagedomain.NewMedicalRecord(
		medicalReEntity.GetID(),
		medicalReEntity.GetAppointmentId(),
		medicalReEntity.GetNursingId(),
		medicalReEntity.GetNursingReport(),
		medicalReEntity.GetStaffConfirm(),
		medicalReEntity.GetStatus(),
		medicalReEntity.GetCreatedAt(),
	)

	if err = h.medicalRecordFetcher.UpdateMedicalRecord(ctx, updateMR); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update medical record when staff assign new nursing to the appointment").
			WithInner(err.Error())
	}

	// Commit transaction if all services created successfully
	if err = h.txManager.Commit(ctx); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot commit transaction").
			WithInner(err.Error())
	}
	return nil
}

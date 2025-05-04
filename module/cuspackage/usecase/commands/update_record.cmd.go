package cuspackagecommands

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/google/uuid"
)

type updateMedicalRecordHanlder struct {
	cmdRepo     CusPackageCommandRepo
	appsFetcher AppointmentFetcher
	txManager   common.TransactionManager
}

func NewUpdateMedicalRecordHandler(cmdRepo CusPackageCommandRepo, appsFetcher AppointmentFetcher, txManager common.TransactionManager) *updateMedicalRecordHanlder {
	return &updateMedicalRecordHanlder{
		cmdRepo:     cmdRepo,
		appsFetcher: appsFetcher,
		txManager:   txManager,
	}
}

func (h *updateMedicalRecordHanlder) Handle(ctx context.Context, dto UpdateMedicalRecordDTO, entity *cuspackagedomain.MedicalRecord) error {
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

	if dto.NursingReport == nil && dto.StaffConfirmation == nil && *dto.NursingReport == "" && *dto.StaffConfirmation == "" {
		return common.NewBadRequestError().WithReason("cannot update medical record when the required information is incomplete")
	}

	if err := h.verifyAllTaskhaveDoneBeforeCanUpdateMedicalRecord(ctx, entity.GetAppointmentId()); err != nil {
		return err
	}

	requester, ok := ctx.Value(common.KeyRequester).(common.Requester)
	if !ok {
		return common.NewUnauthorizedError()
	}
	role := requester.Role()
	sub := requester.UserId()

	curNursingReport := entity.GetNursingReport()
	curStaffConfirm := entity.GetStaffConfirm()

	updateNursingReport := ""
	updateStaffConfirm := ""
	updateStatus := entity.GetStatus()

	if dto.NursingReport != nil && *dto.NursingReport != "" {
		if *entity.GetNursingId() != sub {
			return common.NewBadRequestError().WithReason("only the nurse who performed this service is allowed to submit a nursing's report")
		}
		updateNursingReport = *dto.NursingReport
		updateStaffConfirm = curStaffConfirm
	}

	if dto.StaffConfirmation != nil && *dto.StaffConfirmation != "" {
		if role != "staff" {
			return common.NewBadRequestError().WithReason("only the staff managing this service is allowed to respond to the nurse's report")
		}
		if curNursingReport == "" {
			return common.NewBadRequestError().WithReason("staff cannot respond to the report until the responsible nurse has submitted it")
		}
		updateNursingReport = curNursingReport
		updateStaffConfirm = *dto.StaffConfirmation
		updateStatus = cuspackagedomain.RecordStatusDone
	}

	updateEntity, _ := cuspackagedomain.NewMedicalRecord(
		entity.GetID(),
		entity.GetAppointmentId(),
		entity.GetNursingId(),
		updateNursingReport,
		updateStaffConfirm,
		updateStatus,
		entity.GetCreatedAt(),
	)

	if err := h.cmdRepo.UpdateMedicalRecord(ctx, updateEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update medical record").
			WithInner(err.Error())
	}

	if dto.StaffConfirmation != nil && *dto.StaffConfirmation != "" {
		curAppEntity, err := h.appsFetcher.FindById(ctx, entity.GetAppointmentId())
		if err != nil {
			return common.NewInternalServerError().
				WithReason("cannot get appointment information to update status success").
				WithInner(err.Error())
		}

		newAppEntity, _ := appointmentdomain.NewAppointment(
			curAppEntity.GetID(),
			curAppEntity.GetServiceID(),
			curAppEntity.GetSvcpackageID(),
			curAppEntity.GetCusPackageID(),
			curAppEntity.GetPatientID(),
			curAppEntity.GetNursingID(),
			curAppEntity.GetPatientAddress(),
			curAppEntity.GetPatientLatLng(),
			appointmentdomain.AppStatusSuccess,
			curAppEntity.GetTotalEstDuration(),
			curAppEntity.GetEstDate(),
			curAppEntity.GetActDate(),
			curAppEntity.GetCreatedAt(),
		)

		if err := h.appsFetcher.UpdateAppointment(ctx, newAppEntity); err != nil {
			return common.NewInternalServerError().
				WithReason("cannot update appointment status success").
				WithInner(err.Error())
		}
	}

	// Commit transaction if all services created successfully
	if err = h.txManager.Commit(ctx); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot commit transaction").
			WithInner(err.Error())
	}

	return nil
}

func (h *updateMedicalRecordHanlder) verifyAllTaskhaveDoneBeforeCanUpdateMedicalRecord(ctx context.Context, appsId uuid.UUID) error {
	apps, err := h.appsFetcher.FindById(ctx, appsId)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot get appointment information").
			WithInner(err.Error())
	}

	if err := h.cmdRepo.VerifyAllCusTasksHaveDone(ctx, apps.GetCusPackageID(), apps.GetEstDate()); err != nil {
		if errors.Is(err, common.ErrCustaskNotDoneAll) {
			return common.ErrCustaskNotDoneAll
		} else {
			return common.NewInternalServerError().
				WithReason("cannot get verify all custask have done in this appointment").
				WithInner(err.Error())
		}
	}

	return nil
}

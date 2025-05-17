package cuspackagecommands

import (
	"context"
	"fmt"
	"log"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type cancelPackageHanlder struct {
	cmdRepo         CusPackageCommandRepo
	appsFetcher     AppointmentFetcher
	txManager       common.TransactionManager
	pushNotiFetcher ExternalPushNotiService
}

func NewCancelPackageHandler(
	cmdRepo CusPackageCommandRepo,
	appsFetcher AppointmentFetcher,
	txManager common.TransactionManager,
	pushNotiFetcher ExternalPushNotiService,
) *cancelPackageHanlder {
	return &cancelPackageHanlder{
		cmdRepo:         cmdRepo,
		appsFetcher:     appsFetcher,
		txManager:       txManager,
		pushNotiFetcher: pushNotiFetcher,
	}
}

func (h *cancelPackageHanlder) Handle(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error {
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

	if entity.GetPaymentStatus() == cuspackagedomain.PaymentStatusPaid {
		return common.NewInternalServerError().
			WithReason("cannot cancel the package because this package has already paid")
	}

	apps, err := h.appsFetcher.GetAppointmentByCuspackage(ctx, entity.GetID())
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot get list appointment in this customized package").
			WithInner(err.Error())
	}

	newCuspackage, _ := cuspackagedomain.NewCustomizedPackage(
		entity.GetID(),
		entity.GetServicePackageID(),
		entity.GetPatientID(),
		entity.GetName(),
		entity.GetTotalFee(),
		entity.GetPaidAmount(),
		entity.GetUnpaidAmount(),
		entity.GetPaymentStatus(),
		true,
		entity.GetCreatedAt(),
	)

	if err := h.cmdRepo.UpdateCustomizedPackage(ctx, newCuspackage); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update customized package").
			WithInner(err.Error())
	}

	mappingNurse := make([]DateNursingMapping, len(apps))
	for _, app := range apps {
		newApp, _ := appointmentdomain.NewAppointment(
			app.GetID(),
			app.GetServiceID(),
			app.GetCusPackageID(),
			app.GetPatientID(),
			app.GetNursingID(),
			app.GetPatientAddress(),
			app.GetPatientLatLng(),
			appointmentdomain.AppStatusCancel,
			app.GetTotalEstDuration(),
			app.GetEstDate(),
			app.GetActDate(),
			app.GetCreatedAt(),
		)
		mappingNurse = append(mappingNurse, DateNursingMapping{
			Date:      app.GetEstDate(),
			NursingId: app.GetNursingID(),
		})

		if err := h.appsFetcher.UpdateAppointment(ctx, newApp); err != nil {
			return common.NewInternalServerError().
				WithReason("cannot update appointment").
				WithInner(err.Error())
		}
	}

	// Commit transaction if all services created successfully
	if err = h.txManager.Commit(ctx); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot commit transaction").
			WithInner(err.Error())
	}

	h.PushNotiToNursing(ctx, mappingNurse)

	return nil
}

func (h *cancelPackageHanlder) PushNotiToNursing(ctx context.Context, mappings []DateNursingMapping) {
	for _, obj := range mappings {
		if obj.NursingId != nil {
			contentVi := fmt.Sprintf(
				"Cuộc hẹn dịch vụ vào lúc %s đã bịhuỷ.\n",
				obj.Date.Format("15:04 ngày 02 tháng 01 năm 2006"),
			)
			reqPushNoti := common.PushNotiRequest{
				AccountID: *obj.NursingId,
				Content:   contentVi,
				Route:     "/(tabs)/schedule",
			}
			err_noti := h.pushNotiFetcher.PushNotification(ctx, &reqPushNoti)
			if err_noti != nil {
				log.Println("error push noti for nursing: ", err_noti)
			}
		}
	}
}

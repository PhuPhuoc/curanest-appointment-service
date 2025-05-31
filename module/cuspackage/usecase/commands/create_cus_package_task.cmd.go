package cuspackagecommands

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/payOSHQ/payos-lib-golang"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type createCusPackageAndTaskHandler struct {
	cmdRepo            CusPackageCommandRepo
	svcPackageFetcher  SvcPackageFetcher
	appointmentFetcher AppointmentFetcher
	invoiceFetcher     InvoiceFetcher
	txManager          common.TransactionManager
	payosConfig        common.PayOSConfig
	goongapi           ExternalGoongAPI
	pushNotiFetcher    ExternalPushNotiService
}

func NewCreateCusPackageAndTaskHandler(
	cmdRepo CusPackageCommandRepo,
	svcPackageFetcher SvcPackageFetcher,
	appointmentFetcher AppointmentFetcher,
	invoiceFetcher InvoiceFetcher,
	txManager common.TransactionManager,
	payosConfig common.PayOSConfig,
	goongAPI ExternalGoongAPI,
	pushNotiFetcher ExternalPushNotiService,
) *createCusPackageAndTaskHandler {
	return &createCusPackageAndTaskHandler{
		cmdRepo:            cmdRepo,
		svcPackageFetcher:  svcPackageFetcher,
		appointmentFetcher: appointmentFetcher,
		invoiceFetcher:     invoiceFetcher,
		txManager:          txManager,
		payosConfig:        payosConfig,
		goongapi:           goongAPI,
		pushNotiFetcher:    pushNotiFetcher,
	}
}

func (h *createCusPackageAndTaskHandler) Handle(ctx context.Context, req *ReqCreatePackageTaskDTO) (*uuid.UUID, error) {
	ctx, err := h.txManager.Begin(ctx)
	if err != nil {
		return nil, common.NewInternalServerError().
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

	dateNurseMappings := req.DateNurseMappings
	customizedTasks := req.TaskInfos

	// verify tasks len
	if err = verifyLenOfTask(customizedTasks); err != nil {
		return nil, err
	}

	// get service package to create and verify customized package
	servicePackage, err := h.fetchServicePackage(ctx, req.SvcPackageId)
	if err != nil {
		return nil, err
	}
	// generate uuid of customized-package to create customized-task
	cusPackageId := common.GenUUID()

	// verify the valid of dates
	if err := verifyDates(servicePackage.GetComboDays(), servicePackage.GetTimeInterVal(), dateNurseMappings); err != nil {
		return nil, err
	}

	// if err := h.verifyNurseAvailabilityWithDate(ctx, req.DateNurseMappings); err != nil {
	// 	return nil, err
	// }

	// get list service task of service package above -> to verify customized task before create them
	serviceTasks, err := h.fetchServiceTasks(ctx, servicePackage.GetID())
	if err != nil {
		return nil, err
	}

	cusTaskEnties, totalFee, totalEstDuration, err := validateCustomizedTasks(cusPackageId, serviceTasks, customizedTasks, dateNurseMappings)
	if err != nil {
		return nil, err
	}

	var totalAfterDiscount float64
	discount := servicePackage.GetDiscount()
	if discount == 0 {
		totalAfterDiscount = totalFee
	} else {
		totalAfterDiscount = totalFee * float64(100-discount) / 100
	}

	// create complete dto for creating entity
	cusPackageEntity, _ := cuspackagedomain.NewCustomizedPackage(
		cusPackageId,
		servicePackage.GetID(),
		req.PatientId,
		servicePackage.GetName(),
		totalAfterDiscount,
		0,
		totalAfterDiscount,
		cuspackagedomain.PaymentStatusUnpaid,
		false,
		nil,
	)

	if err = h.savePackageAndTasks(ctx, cusPackageEntity, cusTaskEnties); err != nil {
		return nil, err
	}

	// create appointment
	if err = h.saveAppointment(ctx, servicePackage.GetID(), dateNurseMappings, totalEstDuration, servicePackage.GetServiceID(), req.PatientId, req.PatientAddress, cusPackageEntity); err != nil {
		return nil, err
	}

	// create invoice
	if err = h.saveInvoice(ctx, cusPackageEntity.GetID(), cusPackageEntity.GetTotalFee()); err != nil {
		return nil, err
	}

	// Commit transaction if all services created successfully
	if err = h.txManager.Commit(ctx); err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot commit transaction").
			WithInner(err.Error())
	}

	// h.PushNotiToNursing(ctx, dateNurseMappings)

	objId := cusPackageEntity.GetID()
	return &objId, nil
}

func (h *createCusPackageAndTaskHandler) PushNotiToNursing(ctx context.Context, mappings []DateNursingMapping) {
	for _, obj := range mappings {
		if obj.NursingId != nil {
			vnTime := obj.Date.Add(7 * time.Hour)
			contentVi := fmt.Sprintf(
				"Bạn có cuộc hẹn dịch vụ mới được lên lịch vào lúc %s.\n",
				vnTime.Format("15:04 ngày 02 tháng 01 năm 2006"),
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

func (h *createCusPackageAndTaskHandler) savePackageAndTasks(
	ctx context.Context,
	packageEntity *cuspackagedomain.CustomizedPackage,
	taskEnties []cuspackagedomain.CustomizedTask,
) error {
	// create customized package after verify
	if err := h.cmdRepo.CreateCustomizedPackage(ctx, packageEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create customized-package").
			WithInner(err.Error())
	}

	if err := h.cmdRepo.CreateCustomizedTasks(ctx, taskEnties); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create customized-tasks").
			WithInner(err.Error())
	}

	return nil
}

func (h *createCusPackageAndTaskHandler) getGeocodeFromGoong(ctx context.Context, address string) (string, error) {
	resp, err := h.goongapi.GetGeocodeFromGoong(ctx, address)
	if err != nil {
		return "", err
	}

	if len(resp.Results) == 0 {
		return "", fmt.Errorf("no results found")
	}
	lat := resp.Results[0].Geometry.Location.Lat
	lng := resp.Results[0].Geometry.Location.Lng

	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	geocode := latStr + "," + lngStr

	return geocode, nil
}

func (h *createCusPackageAndTaskHandler) saveAppointment(ctx context.Context, svcpackageId uuid.UUID, mappings []DateNursingMapping, totalEstDuration int, serviceId uuid.UUID, patientId uuid.UUID, patientAddress string, cusPackageEntity *cuspackagedomain.CustomizedPackage) error {
	patientLatLng, err := h.getGeocodeFromGoong(ctx, patientAddress)
	if err != nil {
		log.Println("cannot get geocode with address (" + patientAddress + ") error: " + err.Error())
	}

	appointmentEnties := make([]appointmentdomain.Appointment, len(mappings))
	recordEnties := make([]cuspackagedomain.MedicalRecord, len(mappings))

	for i, obj := range mappings {
		appStatus := appointmentdomain.AppStatusWaiting
		if obj.NursingId != nil {
			appStatus = appointmentdomain.AppStatusConfirmed
		}
		appointmentId := common.GenUUID()

		appointmentEntity, _ := appointmentdomain.NewAppointment(
			appointmentId,
			serviceId,
			cusPackageEntity.GetID(),
			patientId,
			obj.NursingId,
			patientAddress,
			patientLatLng,
			appStatus,
			totalEstDuration,
			obj.Date,
			nil,
			nil,
		)
		appointmentEnties[i] = *appointmentEntity

		recordEntity, _ := cuspackagedomain.NewMedicalRecord(
			common.GenUUID(),
			appointmentId,
			obj.NursingId,
			"",
			"",
			cuspackagedomain.RecordStatusNotDone,
			nil,
		)
		recordEnties[i] = *recordEntity
	}

	if err := h.appointmentFetcher.CreateAppointments(ctx, appointmentEnties); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create appointments").
			WithInner(err.Error())
	}

	if err := h.cmdRepo.CreateMedicalRecords(ctx, recordEnties); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create medical record").
			WithInner(err.Error())
	}

	for _, dto := range appointmentEnties {
		h.PushNotiToNurse(ctx, dto.GetID(), dto.GetNursingID(), dto.GetEstDate())
	}

	return nil
}

func (h *createCusPackageAndTaskHandler) PushNotiToNurse(ctx context.Context, appId uuid.UUID, nursingId *uuid.UUID, appDate time.Time) {
	if nursingId != nil {
		vnTime := appDate.Add(7 * time.Hour)
		contentVi := fmt.Sprintf(
			"Bạn có cuộc hẹn dịch vụ mới được lên lịch vào lúc %s.\n",
			vnTime.Format("15:04 ngày 02 tháng 01 năm 2006"),
		)
		reqPushNoti := common.PushNotiRequest{
			AccountID: *nursingId,
			Content:   contentVi,
			SubID:     appId,
			Route:     "/detail-appointment/[id]",
		}
		err_noti := h.pushNotiFetcher.PushNotification(ctx, &reqPushNoti)
		if err_noti != nil {
			log.Println("error push noti for nursing: ", err_noti)
		}
	}
}

func (h *createCusPackageAndTaskHandler) saveInvoice(ctx context.Context, cusPackageId uuid.UUID, totalFee float64) error {
	invoiceId := common.GenUUID()
	orderCode := time.Now().Unix()*1000 + int64(rand.Intn(1000))
	payos.Key(h.payosConfig.ClientId, h.payosConfig.ApiKey, h.payosConfig.CheckSumKey)

	paymentRequest := payos.CheckoutRequestType{
		OrderCode:   orderCode,
		Amount:      int(totalFee),
		Description: "curanest",
		CancelUrl:   "https://curanest.com.vn/payment-result-fail",
		ReturnUrl:   "https://curanest.com.vn/payment-result-success",
	}

	response, err := payos.CreatePaymentLink(paymentRequest)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("failed to create payment url").
			WithInner(err.Error())
	}

	entity, _ := invoicedomain.NewInvoice(
		invoiceId,
		cusPackageId,
		&orderCode,
		totalFee,
		invoicedomain.PaymentStatusUnpaid,
		"",
		&response.CheckoutUrl,
		&response.QRCode,
		nil,
	)
	if err := h.invoiceFetcher.CreateInvoice(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create invoice for this service").
			WithInner(err.Error())
	}

	return nil
}

func verifyLenOfTask(cusTasks []CreateCustomizedTaskDTO) error {
	if len(cusTasks) == 0 {
		return common.NewBadRequestError().
			WithReason("cannot create services and appoinment without any task")
	}

	return nil
}

func (h *createCusPackageAndTaskHandler) verifyNurseAvailabilityWithDate(ctx context.Context, req []DateNursingMapping) error {
	listUUIDs := make([]uuid.UUID, len(req))
	listDates := make([]time.Time, len(req))

	for i, obj := range req {
		if obj.NursingId != nil {
			listUUIDs[i] = *obj.NursingId
		}
		listDates[i] = obj.Date
	}

	if err := h.appointmentFetcher.AreNursesAvailable(ctx, listUUIDs, listDates); err != nil {
		return common.NewBadRequestError().
			WithReason("one or more nurses are unavailable on one or more dates in the request")
	}
	return nil
}

func verifyDates(comboDays, timeInterval int, mappings []DateNursingMapping) error {
	if len(mappings) != comboDays {
		return common.NewBadRequestError().
			WithReason("the number of dates does not equal the number of combo-days specified in the service")
	}

	for i := 1; i < len(mappings); i++ {
		// calculate the distance between the current date and the previous date (in days)
		daysDiff := int(mappings[i].Date.Truncate(24*time.Hour).Sub(mappings[i-1].Date.Truncate(24*time.Hour)).Hours() / 24)
		if daysDiff < timeInterval {
			mess := fmt.Sprintf("the gap between date %s and %s is %d days, which is less than the required interval of %d days",
				mappings[i-1].Date.Format("2006-01-02"), mappings[i].Date.Format("2006-01-02"), daysDiff, timeInterval)
			return common.NewBadRequestError().WithReason(mess)
		}
	}

	return nil
}

func (h *createCusPackageAndTaskHandler) fetchServicePackage(ctx context.Context, svcPackageId uuid.UUID) (*svcpackagedomain.ServicePackage, error) {
	svcPackage, err := h.svcPackageFetcher.GetServicePackageById(ctx, svcPackageId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get service-package information").
			WithInner(err.Error())
	}
	return svcPackage, nil
}

func (h *createCusPackageAndTaskHandler) fetchServiceTasks(ctx context.Context, svcPackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error) {
	svcTasks, err := h.svcPackageFetcher.GetServiceTasksByPackageId(ctx, svcPackageId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list service tasks").
			WithInner(err.Error())
	}
	return svcTasks, nil
}

func validateCustomizedTasks(cusPackageId uuid.UUID, svcTask []svcpackagedomain.ServiceTask, cusTask []CreateCustomizedTaskDTO, mappings []DateNursingMapping) ([]cuspackagedomain.CustomizedTask, float64, int, error) {
	// create custask map to compare with service task and verify all field before create customized task
	cusTaskMap := make(map[uuid.UUID]CreateCustomizedTaskDTO)
	for _, item := range cusTask {
		cusTaskMap[item.SvcTaskId] = item
	}

	svcTaskMap := make(map[uuid.UUID]svcpackagedomain.ServiceTask)

	// the for loop to check the elements in the svctask array to see if any mandatory task that must exist in the service is missing
	for _, item := range svcTask {
		svcTaskMap[item.GetID()] = item
		if _, existed := cusTaskMap[item.GetID()]; !existed {
			if item.GetIsMustHave() {
				mess := "task (with id: " + item.GetID().String() + " ) must be included in this service"
				return []cuspackagedomain.CustomizedTask{}, 0, 0, common.NewBadRequestError().WithReason(mess)
			}
		}
	}

	// the for loop to check the elements in the svctask array to see if there is any extra task that does not belong to the service
	for _, item := range cusTask {
		if _, existed := svcTaskMap[item.SvcTaskId]; !existed {
			mess := "task (with id: " + item.SvcTaskId.String() + " ) is not included in this service"
			return []cuspackagedomain.CustomizedTask{}, 0, 0, common.NewBadRequestError().WithReason(mess)
		}
	}

	// ***
	// *** verify customized package is valid and handle package combo
	// ***

	// total fee of the service
	var total float64
	// total duration of the service
	totalEstDuration := 0

	// after verify custask from request body -> change dto to entity(domain)
	cusTaskEnties := []cuspackagedomain.CustomizedTask{}
	for i := range mappings {
		for _, item := range cusTask {
			svctask := svcTaskMap[item.SvcTaskId]
			custask, _ := cuspackagedomain.NewCustomizedTask(
				common.GenUUID(),
				item.SvcTaskId,
				cusPackageId,
				svctask.GetTaskOrder(),
				svctask.GetName(),
				item.ClientNote,
				svctask.GetStaffAdvice(),
				item.EstDuration,
				item.TotalCost,
				cuspackagedomain.EnumCusTaskUnit(svctask.GetUnit().String()),
				item.TotalUnit,
				mappings[i].Date,
				nil,
				cuspackagedomain.CusTaskStatusNotDone,
			)
			cusTaskEnties = append(cusTaskEnties, *custask)
			total += item.TotalCost
			if i == 0 {
				totalEstDuration += item.EstDuration
			}
		}
	}

	return cusTaskEnties, total, totalEstDuration, nil
}

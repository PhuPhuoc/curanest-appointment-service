package cuspackagecommands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

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
}

func NewCreateCusPackageAndTaskHandler(
	cmdRepo CusPackageCommandRepo,
	svcPackageFetcher SvcPackageFetcher,
	appointmentFetcher AppointmentFetcher,
	invoiceFetcher InvoiceFetcher,
	txManager common.TransactionManager,
) *createCusPackageAndTaskHandler {
	return &createCusPackageAndTaskHandler{
		cmdRepo:            cmdRepo,
		svcPackageFetcher:  svcPackageFetcher,
		appointmentFetcher: appointmentFetcher,
		invoiceFetcher:     invoiceFetcher,
		txManager:          txManager,
	}
}

func (h *createCusPackageAndTaskHandler) Handle(ctx context.Context, req *ReqCreatePackageTaskDTO) error {
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

	dates := req.Dates
	customizedPackage := req.PackageInfo
	customizedTasks := req.TaskInfos

	// verify tasks len
	if err = verifyLenOfTask(customizedTasks); err != nil {
		return err
	}

	// get service package to create and verify customized package
	servicePackage, err := h.fetchServicePackage(ctx, customizedPackage.SvcPackageId)
	if err != nil {
		return err
	}
	// generate uuid of customized-package to create customized-task
	cusPackageId := common.GenUUID()

	// verify the valid of dates
	if err = verifyDates(servicePackage.GetComboDays(), servicePackage.GetTimeInterVal(), dates); err != nil {
		return err
	}

	// get list service task of service package above -> to verify customized task before create them
	serviceTasks, err := h.fetchServiceTasks(ctx, servicePackage.GetID())
	if err != nil {
		return err
	}

	cusTaskEnties, totalFee, err := validateCustomizedTasks(cusPackageId, serviceTasks, customizedTasks, dates)
	if err != nil {
		return err
	}

	// create complete dto for creating entity
	cusPackageEntity, _ := cuspackagedomain.NewCustomizedPackage(
		cusPackageId,
		servicePackage.GetID(),
		customizedPackage.PatientId,
		servicePackage.GetName(),
		totalFee,
		0,
		totalFee,
		cuspackagedomain.PaymentStatusUnpaid,
		nil,
	)

	if err = h.savePackageAndTasks(ctx, cusPackageEntity, cusTaskEnties); err != nil {
		return err
	}

	// create appointment
	if err = h.saveAppointment(ctx, dates, req.NursingId, req.PatientId, cusPackageEntity); err != nil {
		return err
	}

	// create invoice
	if err = h.saveInvoice(ctx, cusPackageEntity.GetID(), cusPackageEntity.GetTotalFee()); err != nil {
		return err
	}

	// Commit transaction if all services created successfully
	if err = h.txManager.Commit(ctx); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot commit transaction").
			WithInner(err.Error())
	}

	return nil
}

func (h *createCusPackageAndTaskHandler) savePackageAndTasks(ctx context.Context, packageEntity *cuspackagedomain.CustomizedPackage, taskEnties []cuspackagedomain.CustomizedTask) error {
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

func (h *createCusPackageAndTaskHandler) saveAppointment(ctx context.Context, dates []time.Time, nursingId *uuid.UUID, patientId uuid.UUID, cusPackageEntity *cuspackagedomain.CustomizedPackage) error {
	appointmentEnties := make([]appointmentdomain.Appointment, len(dates))
	for i, date := range dates {
		appointmentId := common.GenUUID()
		appointmentEntity, _ := appointmentdomain.NewAppointment(
			appointmentId,
			cusPackageEntity.GetServicePackageID(),
			cusPackageEntity.GetID(),
			patientId,
			nursingId,
			appointmentdomain.AppStatusWaiting,
			date,
			nil,
			nil,
		)
		appointmentEnties[i] = *appointmentEntity
	}

	if err := h.appointmentFetcher.CreateAppointments(ctx, appointmentEnties); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create appointments").
			WithInner(err.Error())
	}
	return nil
}

func (h *createCusPackageAndTaskHandler) saveInvoice(ctx context.Context, cusPackageId uuid.UUID, totalFee float64) error {
	invoiceId := common.GenUUID()
	entity, _ := invoicedomain.NewInvoice(
		invoiceId,
		cusPackageId,
		totalFee,
		invoicedomain.PaymentStatusUnpaid,
		invoicedomain.PaymentTypeWallet,
		"",
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

func verifyDates(comboDays, timeInterval int, dates []time.Time) error {
	if len(dates) != comboDays {
		return common.NewBadRequestError().
			WithReason("the number of dates does not equal the number of combo-days specified in the service")
	}

	for i := 1; i < len(dates); i++ {
		// calculate the distance between the current date and the previous date (in days)
		daysDiff := int(dates[i].Sub(dates[i-1]).Hours() / 24)
		if daysDiff < timeInterval {
			mess := fmt.Sprintf("the gap between date %s and %s is %d days, which is less than the required interval of %d days",
				dates[i-1].Format("2006-01-02"), dates[i].Format("2006-01-02"), daysDiff, timeInterval)
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

func validateCustomizedTasks(cusPackageId uuid.UUID, svcTask []svcpackagedomain.ServiceTask, cusTask []CreateCustomizedTaskDTO, dates []time.Time) ([]cuspackagedomain.CustomizedTask, float64, error) {
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
				return []cuspackagedomain.CustomizedTask{}, 0, common.NewBadRequestError().WithReason(mess)
			}
		}
	}

	// the for loop to check the elements in the svctask array to see if there is any extra task that does not belong to the service
	for _, item := range cusTask {
		if _, existed := svcTaskMap[item.SvcTaskId]; !existed {
			mess := "task (with id: " + item.SvcTaskId.String() + " ) is not included in this service"
			return []cuspackagedomain.CustomizedTask{}, 0, common.NewBadRequestError().WithReason(mess)
		}
	}

	// ***
	// *** verify customized package is valid and handle package combo
	// ***

	// total fee of the service
	var total float64
	// after verify custask from request body -> change dto to entity(domain)
	cusTaskEnties := make([]cuspackagedomain.CustomizedTask, len(cusTask))
	for _, estDate := range dates {
		for i, item := range cusTask {
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
				estDate,
				nil,
				cuspackagedomain.CusTaskStatusNotDone,
			)
			cusTaskEnties[i] = *custask
			total += item.TotalCost
		}
	}

	return cusTaskEnties, total, nil
}

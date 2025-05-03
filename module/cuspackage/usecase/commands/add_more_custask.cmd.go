package cuspackagecommands

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
	"github.com/payOSHQ/payos-lib-golang"
)

type addMoreTaskToAppointmentHandler struct {
	txManager          common.TransactionManager
	cmdRepo            CusPackageCommandRepo
	svcPackageFetcher  SvcPackageFetcher
	appointmentFetcher AppointmentFetcher
	invoiceFetcher     InvoiceFetcher
	payosConfig        common.PayOSConfig
}

func NewAddMoreCusTaskToAppointmentHandler(
	txManager common.TransactionManager,
	cmdRepo CusPackageCommandRepo,
	svcPackageFetcher SvcPackageFetcher,
	appointmentFetcher AppointmentFetcher,
	invoiceFetcher InvoiceFetcher,
	payosConfig common.PayOSConfig,
) *addMoreTaskToAppointmentHandler {
	return &addMoreTaskToAppointmentHandler{
		txManager:          txManager,
		cmdRepo:            cmdRepo,
		svcPackageFetcher:  svcPackageFetcher,
		appointmentFetcher: appointmentFetcher,
		invoiceFetcher:     invoiceFetcher,
		payosConfig:        payosConfig,
	}
}

func (h *addMoreTaskToAppointmentHandler) Handle(ctx context.Context, req *AddMoreCustaskRequestDTO) error {
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
	curApp, err := h.appointmentFetcher.FindById(ctx, req.AppointmentId)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot get appointment").
			WithInner(err.Error())
	}

	svctasks, err := h.svcPackageFetcher.GetServiceTasksByPackageId(ctx, curApp.GetSvcpackageID())
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot get list service-tasks").
			WithInner(err.Error())
	}

	cusTaskEnties, totalFee, totalEstDuration, err := validateAddMoreCustomizedTasks(curApp.GetCusPackageID(), svctasks, req.TaskInfos, curApp.GetEstDate())
	if err != nil {
		return err
	}

	fmt.Println("can phai check lai lich tiep theo cua dieu duong co the se bi trung neu them task nay voi thoi gian du tinh: ", totalEstDuration)

	if err = h.saveInvoice(ctx, curApp.GetCusPackageID(), totalFee); err != nil {
		return err
	}

	if err = h.saveNewCustasks(ctx, cusTaskEnties); err != nil {
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

func (h *addMoreTaskToAppointmentHandler) saveInvoice(ctx context.Context, cusPackageId uuid.UUID, totalFee float64) error {
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
		orderCode,
		totalFee,
		invoicedomain.PaymentStatusUnpaid,
		"",
		response.CheckoutUrl,
		nil,
	)
	if err := h.invoiceFetcher.CreateInvoice(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create invoice for this service").
			WithInner(err.Error())
	}

	return nil
}

func (h *addMoreTaskToAppointmentHandler) saveNewCustasks(
	ctx context.Context,
	taskEnties []cuspackagedomain.CustomizedTask,
) error {
	if err := h.cmdRepo.CreateCustomizedTasks(ctx, taskEnties); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create customized-tasks").
			WithInner(err.Error())
	}

	return nil
}

func validateAddMoreCustomizedTasks(cusPackageId uuid.UUID, svcTask []svcpackagedomain.ServiceTask, cusTask []CreateCustomizedTaskDTO, date time.Time) ([]cuspackagedomain.CustomizedTask, float64, int, error) {
	cusTaskMap := make(map[uuid.UUID]CreateCustomizedTaskDTO)
	for _, item := range cusTask {
		cusTaskMap[item.SvcTaskId] = item
	}

	svcTaskMap := make(map[uuid.UUID]svcpackagedomain.ServiceTask)
	for _, item := range svcTask {
		svcTaskMap[item.GetID()] = item
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
			date,
			nil,
			cuspackagedomain.CusTaskStatusNotDone,
		)
		cusTaskEnties = append(cusTaskEnties, *custask)
		total += item.TotalCost
		totalEstDuration += item.EstDuration
	}

	return cusTaskEnties, total, totalEstDuration, nil
}

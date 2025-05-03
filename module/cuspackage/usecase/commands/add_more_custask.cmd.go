package cuspackagecommands

import (
	"context"
	"fmt"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
)

type addMoreTaskToAppointmentHandler struct {
	cmdRepo            CusPackageCommandRepo
	svcPackageFetcher  SvcPackageFetcher
	appointmentFetcher AppointmentFetcher
	invoiceFetcher     InvoiceFetcher
	payosConfig        common.PayOSConfig
}

func NewAddMoreCusTaskToAppointmentHandler(
	cmdRepo CusPackageCommandRepo,
	svcPackageFetcher SvcPackageFetcher,
	appointmentFetcher AppointmentFetcher,
	invoiceFetcher InvoiceFetcher,
	payosConfig common.PayOSConfig,
) *addMoreTaskToAppointmentHandler {
	return &addMoreTaskToAppointmentHandler{
		cmdRepo:            cmdRepo,
		svcPackageFetcher:  svcPackageFetcher,
		appointmentFetcher: appointmentFetcher,
		invoiceFetcher:     invoiceFetcher,
		payosConfig:        payosConfig,
	}
}

func (h *addMoreTaskToAppointmentHandler) Handle(ctx context.Context, req *AddMoreCustaskRequestDTO) error {
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

	fmt.Println("svctask: ", svctasks)

	return nil
}

func validateAddMoreCustomizedTasks(cusPackageId uuid.UUID, svcTask []svcpackagedomain.ServiceTask, cusTask []CreateCustomizedTaskDTO, date time.Time) ([]cuspackagedomain.CustomizedTask, float64, int, error) {
	// create custask map to compare with service task and verify all field before create customized task
	cusTaskMap := make(map[uuid.UUID]CreateCustomizedTaskDTO)
	for _, item := range cusTask {
		cusTaskMap[item.SvcTaskId] = item
	}

	svcTaskMap := make(map[uuid.UUID]svcpackagedomain.ServiceTask)

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

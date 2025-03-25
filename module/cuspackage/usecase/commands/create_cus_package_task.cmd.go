package cuspackagecommands

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type createCusPackageAndTaskHandler struct {
	cmdRepo           CusPackageCommandRepo
	svcPackageFetcher SvcPackageFetcher
}

func NewCreateCusPackageAndTaskHandler(cmdRepo CusPackageCommandRepo, svcPackageFetcher SvcPackageFetcher) *createCusPackageAndTaskHandler {
	return &createCusPackageAndTaskHandler{
		cmdRepo:           cmdRepo,
		svcPackageFetcher: svcPackageFetcher,
	}
}

func (h *createCusPackageAndTaskHandler) Handle(ctx context.Context, cuspack *CreateCustomizedPackageDTO, custask []CreateCustomizedTaskDTO) error {
	// check custask len if 0 => return err
	if len(custask) == 0 {
		return common.NewBadRequestError().
			WithReason("cannot create services and appoinment without any task")
	}

	// get service package to create and verify customized package
	svcPackage, err := h.svcPackageFetcher.GetServicePackageById(ctx, cuspack.SvcPackageId)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot get service-package information").
			WithInner(err.Error())
	}

	// get list service task of service package above -> to verify customized task before create them
	svcTasks, err := h.svcPackageFetcher.GetServiceTasksByPackageId(ctx, svcPackage.GetID())
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot get list service tasks").
			WithInner(err.Error())
	}

	// create custask map to compare with service task and verify all field before create customized task
	cusTaskMap := make(map[uuid.UUID]CreateCustomizedTaskDTO)
	for _, item := range custask {
		cusTaskMap[item.SvcTaskId] = item
	}

	svcTaskMap := make(map[uuid.UUID]svcpackagedomain.ServiceTask)

	// the for loop to check the elements in the svctask array to see if any mandatory task that must exist in the service is missing
	for _, item := range svcTasks {
		svcTaskMap[item.GetID()] = item
		if _, existed := cusTaskMap[item.GetID()]; !existed {
			if item.GetIsMustHave() {
				mess := "task (with id: " + item.GetID().String() + " ) must be included in this service (with id: " + svcPackage.GetID().String() + " )"
				return common.NewBadRequestError().
					WithReason(mess)
			}
		}
	}

	// the for loop to check the elements in the svctask array to see if there is any extra task that does not belong to the service
	for _, item := range custask {
		if _, existed := svcTaskMap[item.SvcTaskId]; !existed {
			mess := "task (with id: " + item.CusPackageId.String() + " ) is not included in this service (with id: " + svcPackage.GetID().String() + " )"
			return common.NewBadRequestError().
				WithReason(mess)
		}
	}

	// *** verify customized package is valid and handle package combo
	// after verify custask from request body -> change dto to entity(domain)

	cusTaskEnties := make([]cuspackagedomain.CustomizedTask, len(custask))
	for i, item := range custask {
		svctask, existed := svcTaskMap[item.SvcTaskId]
		if !existed {
			mess := "service-task (with id: " + item.CusPackageId.String() + " not found"
			return common.NewBadRequestError().
				WithReason(mess)
		}
		custask, _ := cuspackagedomain.NewCustomizedTask(
			common.GenUUID(),
			item.SvcTaskId,
			item.CusPackageId,
			svctask.GetTaskOrder(),
			svctask.GetName(),
			item.ClientNote,
			svctask.GetStaffAdvice(),
			item.EstDuration,
			item.TotalCost,
			cuspackagedomain.EnumCusTaskUnit(svctask.GetUnit().String()),
			item.TotalUnit,
			item.EstDate,
			nil,
			cuspackagedomain.CusTaskStatusNotDone,
		)
		cusTaskEnties[i] = *custask
	}

	cusPackageId := common.GenUUID()
	cusPackageEntity, _ := cuspackagedomain.NewCustomizedPackage(
		cusPackageId,
		svcPackage.GetID(),
		cuspack.PatientId,
		svcPackage.GetName(),
		cuspack.TotalFee,
		0,
		cuspack.TotalFee,
		cuspackagedomain.PaymentStatusUnpaid,
		nil,
	)

	// create customized package after verify
	if err := h.cmdRepo.CreateCustomizedPackage(ctx, cusPackageEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create customized-package").
			WithInner(err.Error())
	}

	if err := h.cmdRepo.CreateCustomizedTasks(ctx, cusTaskEnties); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create customized-tasks").
			WithInner(err.Error())
	}

	return nil
}

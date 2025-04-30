package cuspackagecommands

import (
	"context"
	"errors"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type updateCustaskStatusDoneHanlder struct {
	cmdRepo     CusPackageCommandRepo
	appsFetcher AppointmentFetcher
}

func NewUpdateCustaskStatusDoneHandler(cmdRepo CusPackageCommandRepo, appsFetcher AppointmentFetcher) *updateCustaskStatusDoneHanlder {
	return &updateCustaskStatusDoneHanlder{
		cmdRepo:     cmdRepo,
		appsFetcher: appsFetcher,
	}
}

func (h *updateCustaskStatusDoneHanlder) Handle(ctx context.Context, entity *cuspackagedomain.CustomizedTask) error {
	if entity.GetStatus() == cuspackagedomain.CusTaskStatusDone {
		return common.NewBadRequestError().
			WithReason("this task is already done")
	}

	if err := h.appsFetcher.CheckAppointmentStatusUpcoming(ctx, entity.GetCusPackageID(), entity.GetEstDate()); err != nil {
		if errors.Is(err, common.ErrAppointmentStatusIsNotUpcoming) {
			return common.NewBadRequestError().
				WithReason("The appointment has not started yet, the task status cannot be updated")
		} else {
			return common.NewInternalServerError().
				WithReason("cannot check appointment status").
				WithInner(err.Error())
		}
	}

	actDate := time.Now()

	updateEntity, _ := cuspackagedomain.NewCustomizedTask(
		entity.GetID(),
		entity.GetSvcTaskID(),
		entity.GetCusPackageID(),
		entity.GetTaskOrder(),
		entity.GetName(),
		entity.GetClientNote(),
		entity.GetStaffAdvice(),
		entity.GetEstDuration(),
		entity.GetTotalCost(),
		entity.GetUnit(),
		entity.GetTotalUnit(),
		entity.GetEstDate(),
		&actDate,
		cuspackagedomain.CusTaskStatusDone,
	)

	if err := h.cmdRepo.UpdateCustomizedTask(ctx, updateEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update custask").
			WithInner(err.Error())
	}

	return nil
}

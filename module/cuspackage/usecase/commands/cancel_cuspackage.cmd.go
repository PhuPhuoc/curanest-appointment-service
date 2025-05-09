package cuspackagecommands

import (
	"context"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type updateCustaskStatusCancelHanlder struct {
	cmdRepo     CusPackageCommandRepo
	appsFetcher AppointmentFetcher
}

func NewUpdateCustaskCancelDoneHandler(cmdRepo CusPackageCommandRepo, appsFetcher AppointmentFetcher) *updateCustaskStatusCancelHanlder {
	return &updateCustaskStatusCancelHanlder{
		cmdRepo:     cmdRepo,
		appsFetcher: appsFetcher,
	}
}

func (h *updateCustaskStatusCancelHanlder) Handle(ctx context.Context, entity *cuspackagedomain.CustomizedTask) error {
	return nil
}

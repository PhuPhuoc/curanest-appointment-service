package cuspackagecommands

import (
	"context"
)

type addMoreTaskToAppointmentHandler struct {
	cmdRepo            CusPackageCommandRepo
	svcPackageFetcher  SvcPackageFetcher
	appointmentFetcher AppointmentFetcher
}

func NewAddMoreCusTaskToAppointmentHandler(
	cmdRepo CusPackageCommandRepo,
	svcPackageFetcher SvcPackageFetcher,
	appointmentFetcher AppointmentFetcher,
) *addMoreTaskToAppointmentHandler {
	return &addMoreTaskToAppointmentHandler{
		cmdRepo:            cmdRepo,
		svcPackageFetcher:  svcPackageFetcher,
		appointmentFetcher: appointmentFetcher,
	}
}

func (h *addMoreTaskToAppointmentHandler) Handle(ctx context.Context, req *AddMoreCustaskRequestDTO) error {
	// curApp, err := h.appointmentFetcher.FindById(ctx, req.AppointmentId)
	// if err != nil {
	// 	return common.NewInternalServerError().
	// 		WithReason("cannot get appointment").
	// 		WithInner(err.Error())
	// }

	return nil
}

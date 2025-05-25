package appointmentqueries

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type getAppointmentsHandler struct {
	queryRepo         AppointmentQueryRepo
	cusPackageFetcher CusPackageFetcher
}

func NewGetAppointmentsHandler(queryRepo AppointmentQueryRepo, cusPackageFetcher CusPackageFetcher) *getAppointmentsHandler {
	return &getAppointmentsHandler{
		queryRepo:         queryRepo,
		cusPackageFetcher: cusPackageFetcher,
	}
}

func (h *getAppointmentsHandler) Handle(ctx context.Context, filter *FilterGetAppointmentDTO) ([]AppointmentDTO, error) {
	if filter.Paging != nil {
		filter.Paging.Process()
	}
	entities, err := h.queryRepo.GetAppointment(ctx, filter)
	if err != nil {
		return []AppointmentDTO{}, common.NewInternalServerError().
			WithReason("cannot get appointment").
			WithInner(err.Error())
	}
	if len(entities) == 0 {
		return []AppointmentDTO{}, nil
	}

	cusPkgIDSet := make(map[uuid.UUID]struct{})
	dtos := make([]AppointmentDTO, len(entities)) // pre-allocate

	for i, entity := range entities {
		dtos[i] = *toAppointmentDTO(&entity)
		if entity.GetCusPackageID() != uuid.Nil {
			cusPkgIDSet[entity.GetCusPackageID()] = struct{}{}
		}
	}

	cusPackIDs := make([]uuid.UUID, 0, len(cusPkgIDSet))
	for id := range cusPkgIDSet {
		cusPackIDs = append(cusPackIDs, id)
	}

	mapCusPackage, err := h.cusPackageFetcher.GetCusPackageByIds(ctx, cusPackIDs)
	if err != nil {
		return []AppointmentDTO{}, common.NewInternalServerError().
			WithReason("cannot get cus-package").
			WithInner(err.Error())
	}

	for i := range dtos {
		cusPackEntity, ok := mapCusPackage[dtos[i].CusPackageId]
		if ok {
			if cusPackEntity.GetPaymentStatus() == cuspackagedomain.PaymentStatusPaid {
				dtos[i].IsPaid = true
			} else {
				dtos[i].IsPaid = false
			}
		}

	}

	return dtos, nil
}

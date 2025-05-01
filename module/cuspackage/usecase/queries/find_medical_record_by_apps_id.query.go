package cuspackagequeries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type findMedicalRecordByAppsIdHandler struct {
	queryRepo CusPackageQueryRepo
}

func NewFindMedicalRecordByAppsIdHandler(queryRepo CusPackageQueryRepo) *findMedicalRecordByAppsIdHandler {
	return &findMedicalRecordByAppsIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findMedicalRecordByAppsIdHandler) Handle(ctx context.Context, appsId uuid.UUID) (*MedicalRecordDTO, error) {
	entity, err := h.queryRepo.FindMedicalRecordByAppsId(ctx, appsId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason(fmt.Sprintf("cannot found medical record with appointment-id: %v", appsId.String())).
			WithInner(err.Error())
	}

	return toMedicalRecordDTO(entity), nil
}

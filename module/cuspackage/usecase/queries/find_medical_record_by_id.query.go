package cuspackagequeries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type findMedicalRecordByIdHandler struct {
	queryRepo CusPackageQueryRepo
}

func NewFindMedicalRecordByIdHandler(queryRepo CusPackageQueryRepo) *findMedicalRecordByIdHandler {
	return &findMedicalRecordByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findMedicalRecordByIdHandler) Handle(ctx context.Context, medicalRecordId uuid.UUID) (*MedicalRecordDTO, error) {
	entity, err := h.queryRepo.FindMedicalRecordById(ctx, medicalRecordId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason(fmt.Sprintf("cannot found medical record with id: %v", medicalRecordId.String())).
			WithInner(err.Error())
	}

	return toMedicalRecordDTO(entity), nil
}

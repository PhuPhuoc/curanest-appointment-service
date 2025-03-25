package cuspackagecommands

import (
	"time"

	"github.com/google/uuid"
)

type CreateCustomizedPackageDTO struct {
	SvcPackageId uuid.UUID `json:"svcpackage-id" binding:"required"`
	PatientId    uuid.UUID `json:"patient-id" binding:"required"`
	TotalFee     float64   `json:"total-fee" binding:"required"`
}

type CreateCustomizedTaskDTO struct {
	SvcTaskId    uuid.UUID `json:"svctask-id" binding:"required"`
	CusPackageId uuid.UUID `json:"cuspackage-id" binding:"required"`
	ClientNote   string    `json:"client-note"`
	EstDuration  int       `json:"est-duration"`
	TotalCost    float64   `json:"total-cost"`
	TotalUnit    int       `json:"total-unit"`
	EstDate      time.Time `json:"est-date"`
}

package servicecommands

import (
	"time"

	"github.com/google/uuid"
)

type CreateServiceDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EstDuration string `json:"est-duration"`
}

type UpdateServiceDTO struct {
	Id          uuid.UUID  `json:"id"`
	CategoryId  uuid.UUID  `json:"category=id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	EstDuration string     `json:"est-duration"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
}

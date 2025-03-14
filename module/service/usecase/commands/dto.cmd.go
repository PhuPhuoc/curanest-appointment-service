package servicecommands

import "github.com/google/uuid"

type CreateServiceDTO struct {
	CategoryId  uuid.UUID `json:"category-id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EstDuration string    `json:"est-duration"`
}

package categorydomain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	id          uuid.UUID
	staffId     *uuid.UUID
	name        string
	description string
	createdAt   *time.Time
}

func (a *Category) GetID() uuid.UUID {
	return a.id
}

func (a *Category) GetStaffId() *uuid.UUID {
	return a.staffId
}

func (a *Category) GetName() string {
	return a.name
}

func (a *Category) GetDescription() string {
	return a.description
}

func (a *Category) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewCategory(id uuid.UUID, staffId *uuid.UUID, name, description string, createdAt *time.Time) (*Category, error) {
	return &Category{
		id:          id,
		staffId:     staffId,
		name:        name,
		description: description,
		createdAt:   createdAt,
	}, nil
}

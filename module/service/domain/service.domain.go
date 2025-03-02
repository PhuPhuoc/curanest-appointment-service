package servicedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	id           uuid.UUID
	categoryId   uuid.UUID
	name         string
	description  string
	thumbnail    string
	est_duration string
	status       Status
	createdAt    *time.Time
}

func (a *Service) GetID() uuid.UUID {
	return a.id
}

func (a *Service) GetCatetgoryID() uuid.UUID {
	return a.categoryId
}

func (a *Service) GetName() string {
	return a.name
}

func (a *Service) GetDescription() string {
	return a.description
}

func (a *Service) GetThumbnail() string {
	return a.thumbnail
}

func (a *Service) GetEstDuration() string {
	return a.est_duration
}

func (a *Service) GetStatus() Status {
	return a.status
}

func (a *Service) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewService(id, categoryId uuid.UUID, name, description, thumbnail, est_duration string, status Status, createdAt *time.Time) (*Service, error) {
	return &Service{
		id:           id,
		categoryId:   categoryId,
		name:         name,
		description:  description,
		thumbnail:    thumbnail,
		est_duration: est_duration,
		status:       status,
		createdAt:    createdAt,
	}, nil
}

type Status int

const (
	StatusAvailable Status = iota
	StatusUnavailable
)

func (r Status) String() string {
	switch r {
	case StatusAvailable:
		return "available"
	case StatusUnavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}

func Enum(s string) Status {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "available":
		return StatusAvailable
	default:
		return StatusUnavailable
	}
}

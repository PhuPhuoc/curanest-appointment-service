package common

import "github.com/google/uuid"

type PushNotiRequest struct {
	AccountID uuid.UUID `json:"account-id"`
	Content   string    `json:"content"`
	SubID     uuid.UUID `json:"sub-id"`
	Route     string    `json:"route"`
}

type PushNotiResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

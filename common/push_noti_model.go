package common

import "github.com/google/uuid"

type PushNotiRequest struct {
	AccountID uuid.UUID `json:"account-id"`
	Content   string    `json:"content"`
	Route     string    `json:"route"`
}

type PushNotiResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

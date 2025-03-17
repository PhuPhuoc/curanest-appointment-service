package svcpackagecommands

import "github.com/google/uuid"

type CreateServicePackageDTO struct {
	ServiceId    uuid.UUID `json:"service-id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	ComboDays    int       `json:"combo-days"`
	Discount     int       `json:"discount"`
	TimeInterval int       `json:"time-interval"`
}

type CreateServiceTaskDTO struct {
	IsMustHave         bool    `json:"is-must-have"`
	Order              int     `json:"order"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	StaffAdvice        string  `json:"staff-advice"`
	EstDuration        int     `json:"est-duration"`
	Cost               float64 `json:"cost" binding:"required"`
	AdditionalCost     float64 `json:"additional-cost" binding:"required"`
	AdditionalCostDesc string  `json:"additional-cost-desc" binding:"required"`
	Unit               string  `json:"unit" binding:"oneof=quantity time"`
	PriceOfStep        int     `json:"price-of-step" binding:"required"`
}

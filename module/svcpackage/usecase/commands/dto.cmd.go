package svcpackagecommands

import "github.com/google/uuid"

type CreateServicePackageDTO struct {
	ServiceId    uuid.UUID `json:"service-id" binding:"required uuid"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	ComboDays    int       `json:"combo-days" binding:"required"`
	Discount     int       `json:"discount" binding:"required"`
	TimeInterval int       `json:"time-interval" binding:"required"`
}

type CreateServiceTaskDTO struct {
	IsMustHave         bool    `json:"is-must-have" binding:"required"`
	Order              int     `json:"order" binding:"required"`
	Name               string  `json:"name" binding:"required"`
	Description        string  `json:"description" binding:"required"`
	StaffAdvice        string  `json:"staff-advice" binding:"required"`
	EstDuration        int     `json:"est-duration" binding:"required"`
	Cost               float64 `json:"cost" binding:"required"`
	AdditionalCost     float64 `json:"additional-cost" binding:"required"`
	AdditionalCostDesc string  `json:"additional-cost-desc" binding:"required"`
	Unit               string  `json:"unit" binding:"required,oneof=quantity time"`
	PriceOfStep        int     `json:"price-of-step" binding:"required"`
}

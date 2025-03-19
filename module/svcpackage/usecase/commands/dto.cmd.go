package svcpackagecommands

import "github.com/google/uuid"

type ServicePackageDTO struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	ComboDays    int    `json:"combo-days"`
	Discount     int    `json:"discount"`
	TimeInterval int    `json:"time-interval"`
}

type UpdateServicePackageDTO struct {
	SvcPackageId uuid.UUID `json:"-"`
	ServiceId    uuid.UUID `json:"-"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	ComboDays    int       `json:"combo-days"`
	Discount     int       `json:"discount"`
	TimeInterval int       `json:"time-interval"`
	Status       string    `json:"status" binding:"oneof=available unavailable"`
}

type ServiceTaskDTO struct {
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

type UpdateServiceTaskDTO struct {
	SvcTaskId          uuid.UUID `json:"-"`
	SvcPackageId       uuid.UUID `json:"-"`
	IsMustHave         bool      `json:"is-must-have"`
	Order              int       `json:"order"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	StaffAdvice        string    `json:"staff-advice"`
	EstDuration        int       `json:"est-duration"`
	Cost               float64   `json:"cost" binding:"required"`
	AdditionalCost     float64   `json:"additional-cost" binding:"required"`
	AdditionalCostDesc string    `json:"additional-cost-desc" binding:"required"`
	Unit               string    `json:"unit" binding:"oneof=quantity time"`
	PriceOfStep        int       `json:"price-of-step" binding:"required"`
	Status             string    `json:"status" binding:"oneof=available unavailable"`
}

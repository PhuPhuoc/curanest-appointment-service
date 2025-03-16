package svcpackagerepository

import (
	"time"

	"github.com/google/uuid"
)

var (
	TABLE_PACKAGE = `service_packages`
	TABLE_TASK    = `service_tasks`

	CREATE_FIELD_PACKAGE = []string{
		"id",
		"service_id",
		"name",
		"description",
		"combo_days",
		"discount",
		"time_interval",
		"status",
	}
	CREATE_FIELD_TASK = []string{
		"id",
		"service_package_id",
		"is_must_have",
		"order",
		"name",
		"description",
		"staff_advice",
		"est_duration",
		"cost",
		"addtional_cost",
		"addtional_cost_desc",
		"unit",
		"price_of_step",
		"status",
	}

	GET_FIELD_PACKAGE = []string{
		"id",
		"service_id",
		"name",
		"description",
		"combo_days",
		"discount",
		"time_interval",
		"status",
		"created_at",
	}
	GET_FIELD_TASK = []string{
		"id",
		"service_package_id",
		"is_must_have",
		"order",
		"name",
		"description",
		"staff_advice",
		"est_duration",
		"cost",
		"addtional_cost",
		"addtional_cost_desc",
		"unit",
		"price_of_step",
		"status",
	}

	UPDATE_FIELD_PACKAGE = []string{
		"name",
		"description",
		"combo_days",
		"discount",
		"time_interval",
		"status",
	}
	UPDATE_FIELD_TASK = []string{
		"is_must_have",
		"order",
		"name",
		"description",
		"staff_advice",
		"est_duration",
		"cost",
		"addtional_cost",
		"addtional_cost_desc",
		"unit",
		"price_of_step",
		"status",
	}
)

type SvcPackageDTO struct {
	Id           uuid.UUID  `db:"id"`
	ServiceId    uuid.UUID  `db:"service_id"`
	Name         string     `db:"name"`
	Description  string     `db:"description"`
	ComboDays    int        `db:"combo_days"`
	Discount     int        `db:"discount"`
	TimeInterval int        `db:"time_interval"`
	Status       string     `db:"status"`
	CreatedAt    *time.Time `db:"created_at"`
}

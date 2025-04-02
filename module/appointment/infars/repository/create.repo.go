package appointmentrepository

import (
	"context"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

func (repo *appointmentRepo) CreateAppointments(ctx context.Context, entities []appointmentdomain.Appointment) error {
	return nil
}

package appointmentrepository

import (
	"context"
	"time"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

func (repo *appointmentRepo) GetAppointmentInDate(ctx context.Context, estStartDate, estEndDate time.Time) ([]appointmentdomain.Appointment, error) {
	query := `
		select id, nursing_id from appointments where
		est_date >= ? and est_date <= ?
	`
	var dtos []AppointmentDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, estStartDate, estEndDate); err != nil {
		return []appointmentdomain.Appointment{}, err
	}

	entities := make([]appointmentdomain.Appointment, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToAppointmentEntity()
		entities[i] = *entity
	}

	return entities, nil
}

package appointmentrepository

import (
	"context"
	"time"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

func (repo *appointmentRepo) GetAppointmentInDate(ctx context.Context, estStartDate, estEndDate time.Time) ([]appointmentdomain.Appointment, error) {
	query := `
		select id, nursing_id from appointments where
		est_date < ? and date_add(est_date, interval total_est_duration minute) > ?
		
	`
	var dtos []AppointmentDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, estEndDate, estStartDate); err != nil {
		return []appointmentdomain.Appointment{}, err
	}

	entities := make([]appointmentdomain.Appointment, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToAppointmentEntity()
		entities[i] = *entity
	}

	return entities, nil
}

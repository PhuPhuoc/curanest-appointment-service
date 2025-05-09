package appointmentrepository

import (
	"context"
	"time"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/google/uuid"
)

func (repo *appointmentRepo) GetAppointmentInADayOfNursing(ctx context.Context, nursingId uuid.UUID, estStartDate, estEndDate time.Time) ([]appointmentdomain.Appointment, error) {
	query := `
		select id, nursing_id, est_date, total_est_duration from appointments where
		nursing_id = ? and est_date >= ? and est_date <= ? and status != 'cancel'
		order by est_date desc
	`
	var dtos []AppointmentDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, nursingId, estStartDate, estEndDate); err != nil {
		return []appointmentdomain.Appointment{}, err
	}

	entities := make([]appointmentdomain.Appointment, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToAppointmentEntity()
		entities[i] = *entity
	}

	return entities, nil
}

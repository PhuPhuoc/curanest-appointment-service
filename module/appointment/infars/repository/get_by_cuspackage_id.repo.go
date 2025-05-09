package appointmentrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/google/uuid"
)

func (repo *appointmentRepo) GetAppointmentByCuspackage(ctx context.Context, cuspackageId uuid.UUID) ([]appointmentdomain.Appointment, error) {
	where := "customized_package_id = ?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE_APPOINTMENT, GET_APPOINTMENT, &where)

	var dtos []AppointmentDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, cuspackageId); err != nil {
		return []appointmentdomain.Appointment{}, err
	}

	entities := make([]appointmentdomain.Appointment, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToAppointmentEntity()
		entities[i] = *entity
	}

	return entities, nil
}

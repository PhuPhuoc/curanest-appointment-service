package appointmentrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/google/uuid"
)

func (repo *appointmentRepo) FindById(ctx context.Context, appointmentId uuid.UUID) (*appointmentdomain.Appointment, error) {
	var dto AppointmentDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE_APPOINTMENT, GET_APPOINTMENT, &where)
	if err := repo.db.Get(&dto, query, appointmentId); err != nil {
		return nil, err
	}
	return dto.ToAppointmentEntity()
}

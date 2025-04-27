package appointmentrepository

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

func (repo *appointmentRepo) UpdateAppointment(ctx context.Context, entity *appointmentdomain.Appointment) error {
	dto := ToAppointmentDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_APPOINTMENT, UPDATE_APPOINTMENT, &where)
	fmt.Println("query: ", query)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}

	return nil
}

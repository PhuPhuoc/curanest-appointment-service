package appointmentrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

func (repo *appointmentRepo) UpdateAppointment(ctx context.Context, entity *appointmentdomain.Appointment) error {
	dto := ToAppointmentDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_APPOINTMENT, UPDATE_APPOINTMENT, &where)
	// Get transaction from context if exist
	if tx := common.GetTxFromContext(ctx); tx != nil {
		_, err := tx.NamedExec(query, dto)
		return err
	}

	// If no transaction, use db directly
	_, err := repo.db.NamedExec(query, dto)
	return err
}

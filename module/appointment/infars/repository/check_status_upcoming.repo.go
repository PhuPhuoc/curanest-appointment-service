package appointmentrepository

import (
	"context"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/google/uuid"
)

func (repo *appointmentRepo) CheckAppointmentStatusUpcoming(ctx context.Context, cuspackageId uuid.UUID, date time.Time) error {
	query := `
		select id, status from appointments where
		customized_package_id = ? and est_date = ? 
	`
	var dto AppointmentDTO
	if err := repo.db.GetContext(ctx, &dto, query, cuspackageId, date); err != nil {
		return err
	}

	if dto.Status != appointmentdomain.AppStatusUpcoming.String() {
		return common.ErrAppointmentStatusIsNotUpcoming
	}
	return nil
}

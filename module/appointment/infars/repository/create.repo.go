package appointmentrepository

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/jmoiron/sqlx"
)

func (repo *appointmentRepo) CreateAppointments(ctx context.Context, entities []appointmentdomain.Appointment) error {
	query := common.GenerateSQLQueries(common.INSERT, TABLE_APPOINTMENT, CREATE_APPOINTMENT, nil)

	// get transaction from context if exist
	if tx := common.GetTxFromContext(ctx); tx != nil {
		stmt, err := tx.PrepareNamedContext(ctx, query)
		if err != nil {
			return fmt.Errorf("[create-appointments] prepare statement failed: %w", err)
		}
		defer stmt.Close()

		for i, entity := range entities {
			dto := ToAppointmentDTO(&entity)
			_, err := stmt.ExecContext(ctx, dto)
			if err != nil {
				return fmt.Errorf("[create-appointments] insert failed at index %d: %w", i, err)
			}
		}
		return nil
	}

	var err error
	var tx *sqlx.Tx
	tx, err = repo.db.Beginx()
	if err != nil {
		return fmt.Errorf("[create-appointments] cannot begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("[create-appointments] cannot commit transaction: %w", commitErr)
		}
	}()

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("[create-appointments] prepare statement failed: %w", err)
	}
	defer stmt.Close()

	for i, entity := range entities {
		dto := ToAppointmentDTO(&entity)
		_, err = stmt.ExecContext(ctx, dto)
		if err != nil {
			return fmt.Errorf("[create-appointments] insert failed at index %d: %w", i, err)
		}
	}

	return nil
}

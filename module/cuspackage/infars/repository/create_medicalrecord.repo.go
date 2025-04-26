package cuspackagerepository

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/jmoiron/sqlx"
)

func (repo *cusPackageRepo) CreateMedicalRecords(ctx context.Context, entities []cuspackagedomain.MedicalRecord) error {
	query := common.GenerateSQLQueries(common.INSERT, TABLE_MEDICALRECORD, CREATE_MEDICALRECORD, nil)

	// get transaction from context if exist
	if tx := common.GetTxFromContext(ctx); tx != nil {
		stmt, err := tx.PrepareNamedContext(ctx, query)
		if err != nil {
			return fmt.Errorf("[create-medical-record] prepare statement failed: %w", err)
		}
		defer stmt.Close()

		for i, entity := range entities {
			dto := ToMedicalRecordDTO(&entity)
			_, err := stmt.ExecContext(ctx, dto)
			if err != nil {
				return fmt.Errorf("[create-medical-record] insert failed at index %d: %w", i, err)
			}
		}
		return nil
	}

	var err error
	var tx *sqlx.Tx
	tx, err = repo.db.Beginx()
	if err != nil {
		return fmt.Errorf("[create-medical-record] cannot begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction: %w", commitErr)
		}
	}()

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("[create-medical-record] prepare statement failed: %w", err)
	}
	defer stmt.Close()

	for i, entity := range entities {
		dto := ToMedicalRecordDTO(&entity)
		_, err = stmt.ExecContext(ctx, dto)
		if err != nil {
			return fmt.Errorf("[create-medical-record] insert failed at index %d: %w", i, err)
		}
	}

	return nil
}

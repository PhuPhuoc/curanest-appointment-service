package cuspackagerepository

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/jmoiron/sqlx"
)

func (repo *cusPackageRepo) CreateCustomizedTasks(ctx context.Context, entities []cuspackagedomain.CustomizedTask) error {
	var err error
	var tx *sqlx.Tx
	tx, err = repo.db.Beginx()
	if err != nil {
		return fmt.Errorf("cannot begin transaction: %w", err)
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

	query := common.GenerateSQLQueries(common.INSERT, TABLE_CUSTASK, CREATE_CUSTASK, nil)
	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare statement failed: %w", err)
	}
	defer stmt.Close()

	for i, entity := range entities {
		dto := ToCusTaskDTO(&entity)
		_, err := stmt.ExecContext(ctx, dto)
		if err != nil {
			return fmt.Errorf("insert failed at index %d: %w", i, err)
		}
	}

	return nil
}

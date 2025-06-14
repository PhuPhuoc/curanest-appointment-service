package cuspackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

func (repo *cusPackageRepo) UpdateCustomizedTask(ctx context.Context, entity *cuspackagedomain.CustomizedTask) error {
	dto := ToCusTaskDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_CUSTASK, UPDATE_TASK, &where)
	// Get transaction from context if exist
	if tx := common.GetTxFromContext(ctx); tx != nil {
		_, err := tx.NamedExec(query, dto)
		return err
	}

	// If no transaction, use db directly
	_, err := repo.db.NamedExec(query, dto)
	return err
}

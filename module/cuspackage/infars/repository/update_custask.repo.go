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
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}

	return nil
}

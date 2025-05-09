package svcpackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

func (repo *svcPackageRepo) UpdateTask(ctx context.Context, entity *svcpackagedomain.ServiceTask) error {
	dto := ToSvcTaskDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_TASK, UPDATE_FIELD_TASK, &where)

	// If no transaction, use db directly
	_, err := repo.db.NamedExec(query, dto)
	return err
}

package svcpackagerepository

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

func (repo *svcPackageRepo) CreateTask(ctx context.Context, entity *svcpackagedomain.ServiceTask) error {
	dto := ToSvcTaskDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE_TASK, CREATE_FIELD_TASK, nil)
	fmt.Println("query: ", query)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}

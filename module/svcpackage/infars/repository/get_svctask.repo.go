package svcpackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
)

func (repo *svcPackageRepo) GetSvcTasks(ctx context.Context, svcpackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error) {
	var args []interface{}
	where := "service_package_id=?"
	order := " ORDER BY task_order"
	args = append(args, svcpackageId.String())
	selectQuery := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_TASK, GET_FIELD_TASK, &where)
	query := selectQuery + order

	var dtos []SvcTaskDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, args...); err != nil {
		return nil, err
	}

	if len(dtos) == 0 {
		return []svcpackagedomain.ServiceTask{}, nil
	}

	entities := make([]svcpackagedomain.ServiceTask, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToSvcTaskEntity()
		entities[i] = *entity
	}

	return entities, nil
}

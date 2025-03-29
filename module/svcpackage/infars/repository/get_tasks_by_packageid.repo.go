package svcpackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
)

func (repo *svcPackageRepo) GetServiceTasksByPackageId(ctx context.Context, svcPackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error) {
	where := "service_package_id=?"
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_TASK, GET_FIELD_TASK, &where)

	var dtos []SvcTaskDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, svcPackageId); err != nil {
		return nil, err
	}

	entities := make([]svcpackagedomain.ServiceTask, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToSvcTaskEntity()
		entities[i] = *entity
	}

	return entities, nil
}

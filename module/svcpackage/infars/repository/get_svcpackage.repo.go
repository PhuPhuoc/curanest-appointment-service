package svcpackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
)

func (repo *svcPackageRepo) GetSvcPackges(ctx context.Context, serviceId uuid.UUID) ([]svcpackagedomain.ServicePackage, error) {
	var args []interface{}
	where := "service_id=?"
	order := " ORDER BY created_at desc"
	args = append(args, serviceId.String())
	selectQuery := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_PACKAGE, GET_FIELD_PACKAGE, &where)
	query := selectQuery + order

	var dtos []SvcPackageDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, args...); err != nil {
		return nil, err
	}

	if len(dtos) == 0 {
		return []svcpackagedomain.ServicePackage{}, nil
	}

	entities := make([]svcpackagedomain.ServicePackage, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToSvcPackageEntity()
		entities[i] = *entity
	}

	return entities, nil
}

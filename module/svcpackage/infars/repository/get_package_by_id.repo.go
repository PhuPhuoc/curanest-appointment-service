package svcpackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
)

func (repo *svcPackageRepo) GetServicePackageById(ctx context.Context, svcPackageId uuid.UUID) (*svcpackagedomain.ServicePackage, error) {
	where := "id=?"
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_PACKAGE, GET_FIELD_PACKAGE, &where)

	var dto SvcPackageDTO
	if err := repo.db.GetContext(ctx, &dto, query, svcPackageId); err != nil {
		return nil, err
	}

	entity, _ := dto.ToSvcPackageEntity()
	return entity, nil
}

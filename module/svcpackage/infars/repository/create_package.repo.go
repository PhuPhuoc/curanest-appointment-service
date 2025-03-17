package svcpackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

func (repo *svcPackageRepo) CreatePackage(ctx context.Context, entity *svcpackagedomain.ServicePackage) error {
	dto := ToSvcPackageDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE_PACKAGE, CREATE_FIELD_PACKAGE, nil)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}

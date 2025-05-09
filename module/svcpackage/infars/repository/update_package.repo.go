package svcpackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

func (repo *svcPackageRepo) UpdatePackage(ctx context.Context, entity *svcpackagedomain.ServicePackage) error {
	dto := ToSvcPackageDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_PACKAGE, UPDATE_FIELD_PACKAGE, &where)

	// If no transaction, use db directly
	_, err := repo.db.NamedExec(query, dto)
	return err
}

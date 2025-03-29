package cuspackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

func (repo *cusPackageRepo) CreateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error {
	dto := ToCusPackageDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE_CUSPACKAGE, CREATE_CUSPACKAGE, nil)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}

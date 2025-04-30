package cuspackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/google/uuid"
)

func (repo *cusPackageRepo) FindCusTaskById(ctx context.Context, custaskId uuid.UUID) (*cuspackagedomain.CustomizedTask, error) {
	var dto CusTaskDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE_CUSTASK, GET_CUSTASK, &where)
	if err := repo.db.Get(&dto, query, custaskId); err != nil {
		return nil, err
	}
	return dto.ToCusTaskEntity()
}

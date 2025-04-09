package cuspackagerepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

func (repo *cusPackageRepo) FindCusPackage(ctx context.Context, id uuid.UUID) (*cuspackagedomain.CustomizedPackage, error) {
	var dto CusPackageDTO
	where := "id = ?"
	query := common.GenerateSQLQueries(common.FIND, TABLE_CUSPACKAGE, GET_CUSPACKAGE, &where)
	if err := repo.db.Get(&dto, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return dto.ToCusPackageEntity()
}

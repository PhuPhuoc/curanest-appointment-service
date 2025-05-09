package categoryrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

func (repo *categoryRepo) Update(ctx context.Context, entity *categorydomain.Category) error {
	where := "id=:id"
	dto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.UPDATE, TABLE, UPDATE_FIELD, &where)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}

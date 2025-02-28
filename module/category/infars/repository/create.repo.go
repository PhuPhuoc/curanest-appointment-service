package categoryrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

func (repo *categoryRepo) Create(ctx context.Context, entity *categorydomain.Category) error {
	accdto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, FIELD, nil)
	if _, err := repo.db.NamedExec(query, accdto); err != nil {
		return err
	}
	return nil
}

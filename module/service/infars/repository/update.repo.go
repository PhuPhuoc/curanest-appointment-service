package servicerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

func (repo *serviceRepo) Update(ctx context.Context, entity *servicedomain.Service) error {
	where := "id=:id"
	dto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.UPDATE, TABLE, CREATE_FIELD, &where)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}

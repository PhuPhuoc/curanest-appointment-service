package servicerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

func (repo *serviceRepo) Create(ctx context.Context, entity *servicedomain.Service) error {
	dto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, CREATE_FIELD, nil)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}

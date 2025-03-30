package servicerepository

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
)

func (repo *serviceRepo) GetServicesByCategoryAndFilter(ctx context.Context, cateId uuid.UUID, filter servicequeries.FilterGetService) ([]servicedomain.Service, error) {
	var args []interface{}
	where := "category_id=? "
	args = append(args, cateId.String())
	if filter.ServiceName != "" {
		where = where + "And name like ?"
		args = append(args, "%"+filter.ServiceName+"%")
	}
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE, GET_FIELD, &where)

	var dtos []ServiceDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, args...); err != nil {
		return nil, err
	}

	entities := make([]servicedomain.Service, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}

	return entities, nil
}

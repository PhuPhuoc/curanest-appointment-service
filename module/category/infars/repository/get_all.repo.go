package categoryrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
	categoryqueries "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/queries"
)

func (repo *categoryRepo) GetCategories(ctx context.Context, filter *categoryqueries.FilterCategoryDTO) ([]categorydomain.Category, error) {
	where := ""
	var args []interface{}
	if filter != nil && filter.Name != "" {
		where = "WHERE name like ?"
		args = append(args, "%"+filter.Name+"%")
	}
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE, FIELD, &where)
	queryWhere := query + where

	var dtos []CategoryDTO
	if err := repo.db.SelectContext(ctx, &dtos, queryWhere, args...); err != nil {
		return nil, err
	}

	entities := make([]categorydomain.Category, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}

	return entities, nil
}

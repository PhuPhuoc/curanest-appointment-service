package categoryrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
	"github.com/google/uuid"
)

func (repo *categoryRepo) FindCategoryById(ctx context.Context, cateId uuid.UUID) (*categorydomain.Category, error) {
	var dto CategoryDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE, FIELD, &where)
	if err := repo.db.Get(&dto, query, cateId); err != nil {
		return nil, err
	}
	return dto.ToEntity()
}

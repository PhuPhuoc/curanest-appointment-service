package categoryrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
	"github.com/google/uuid"
)

func (repo *categoryRepo) GetCategoryOfStaff(ctx context.Context, staffId uuid.UUID) (*categorydomain.Category, error) {
	var dto CategoryDTO
	where := "staff_id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE, FIELD, &where)
	if err := repo.db.Get(&dto, query, staffId); err != nil {
		return nil, err
	}
	return dto.ToEntity()
}

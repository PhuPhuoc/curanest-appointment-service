package cuspackagerepository

import (
	"context"
	"fmt"
	"strings"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/google/uuid"
)

func (repo *cusPackageRepo) GetCusPackageByIds(ctx context.Context, cuspackageId []uuid.UUID) (map[uuid.UUID]cuspackagedomain.CustomizedPackage, error) {
	cuspackageIdStrs := make([]string, len(cuspackageId))
	for i, id := range cuspackageId {
		cuspackageIdStrs[i] = id.String()
	}
	cuspackageIdsParam := "'" + strings.Join(cuspackageIdStrs, "','") + "'"

	where := "id IN (%s)"
	query := common.GenerateSQLQueries(common.FIND, TABLE_CUSPACKAGE, GET_CUSPACKAGE, &where)
	query = fmt.Sprintf(query, cuspackageIdsParam)

	var dtos []CusPackageDTO
	if err := repo.db.SelectContext(ctx, &dtos, query); err != nil {
		return nil, err
	}

	mapCuspack := make(map[uuid.UUID]cuspackagedomain.CustomizedPackage, len(dtos))
	for i, dto := range dtos {
		entity, _ := dtos[i].ToCusPackageEntity()
		mapCuspack[dto.Id] = *entity
	}

	return mapCuspack, nil
}

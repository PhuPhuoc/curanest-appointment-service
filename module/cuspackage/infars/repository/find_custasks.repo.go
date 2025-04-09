package cuspackagerepository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/google/uuid"
)

func (repo *cusPackageRepo) FindCusTasks(ctx context.Context, packageId uuid.UUID, estDate time.Time) ([]cuspackagedomain.CustomizedTask, error) {
	var whereConditions []string
	var args []interface{}

	whereConditions = append(whereConditions, "customized_package_id = ?")
	args = append(args, packageId)
	whereConditions = append(whereConditions, "est_date = ?")
	args = append(args, estDate.Format("2006-01-02 15:04:05"))
	where := strings.Join(whereConditions, " AND ")

	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_CUSTASK, GET_CUSTASK, &where)
	fmt.Println("query: ", query)
	var dtos []CusTaskDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, args...); err != nil {
		return nil, err
	}

	fmt.Println("dtos: ", dtos)
	entities := make([]cuspackagedomain.CustomizedTask, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToCusTaskEntity()
		entities[i] = *entity
	}

	return entities, nil
}

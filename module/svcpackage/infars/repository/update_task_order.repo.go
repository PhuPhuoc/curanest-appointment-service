package svcpackagerepository

import (
	"context"
	"strconv"
	"strings"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

func (repo *svcPackageRepo) UpdateTaskOrder(ctx context.Context, entities []svcpackagedomain.ServiceTask) error {
	if len(entities) == 0 {
		return nil
	}
	var builder strings.Builder
	builder.WriteString("UPDATE " + TABLE_TASK)
	builder.WriteString(" SET task_order = CASE ")
	taskids := []string{}
	for _, entity := range entities {
		builder.WriteString(" WHEN id = '" + entity.GetID().String() + "' THEN " + strconv.Itoa(entity.GetTaskOrder()))
		taskids = append(taskids, entity.GetID().String())
	}
	result := "'" + strings.Join(taskids, "', '") + "'"
	fullResult := "(" + result + ")"
	builder.WriteString(" END WHERE id IN " + fullResult)
	query := builder.String()

	if result, err := repo.db.ExecContext(ctx, query); err != nil {
		rowAffect, _ := result.RowsAffected()
		if rowAffect == 0 {
			return nil
		}
		return err
	}
	return nil
}

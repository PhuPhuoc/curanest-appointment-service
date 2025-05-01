package cuspackagerepository

import (
	"context"
	"fmt"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

func (repo *cusPackageRepo) VerifyAllCusTasksHaveDone(ctx context.Context, cusPackageId uuid.UUID, date time.Time) error {
	query := `
		select exists (
			select 1 from customized_tasks 
			where customized_package_id = ? and est_date = ? and status != 'done'
		) as has_not_done;
	`

	var hasNotDone bool
	if err := repo.db.GetContext(ctx, &hasNotDone, query, cusPackageId, date.Format("2006-01-02 15:04:05")); err != nil {
		return fmt.Errorf("failed to verify customized tasks status: %w", err)
	}

	if hasNotDone {
		return common.ErrCustaskNotDoneAll
	}
	return nil
}

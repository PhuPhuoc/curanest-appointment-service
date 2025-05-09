package appointmentrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

func (repo *appointmentRepo) IsNurseAvailability(ctx context.Context, nursingId uuid.UUID, startDate, endDate time.Time) error {
	query := `
		SELECT EXISTS 
		(
			SELECT id 
			FROM appointments 
			WHERE nursing_id = ? and est_date >= ? and est_date <= ? and status != 'cancel'
		)
	`
	var exists bool
	if err := repo.db.GetContext(ctx, &exists, query, nursingId, startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")); err != nil {
		return fmt.Errorf("failed to check nurse availability: %w", err)
	}

	fmt.Printf("nursing_id: %v \nstartDate :%v \nendDate: %v exists: %v \n", nursingId, startDate, endDate, exists)

	if exists {
		return common.ErrNurseNotAvailable
	}

	return nil
}

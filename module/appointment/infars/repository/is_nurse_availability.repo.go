package appointmentrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

func (repo *appointmentRepo) IsNurseAvailability(ctx context.Context, nursingId uuid.UUID, date time.Time) error {
	query := `
		SELECT EXISTS 
		(
			SELECT id 
			FROM appointments 
			WHERE nursing_id = ? AND est_date = ? 
		)
	`
	var exists bool
	if err := repo.db.GetContext(ctx, &exists, query, nursingId, date.Format("2006-01-02 15:04:05")); err != nil {
		return fmt.Errorf("failed to check nurse availability: %w", err)
	}

	if exists {
		return common.ErrNurseNotAvailable
	}

	return nil
}

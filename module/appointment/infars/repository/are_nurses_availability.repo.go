package appointmentrepository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

func (repo *appointmentRepo) AreNursesAvailable(ctx context.Context, nursingIds []uuid.UUID, dates []time.Time) error {
	// Kiểm tra input rỗng
	if len(nursingIds) == 0 || len(dates) == 0 {
		return fmt.Errorf("nursingIds or dates cannot be empty")
	}

	nursingIdStrs := make([]string, len(nursingIds))
	for i, id := range nursingIds {
		nursingIdStrs[i] = id.String()
	}
	nursingIdsParam := "'" + strings.Join(nursingIdStrs, "','") + "'"

	dateStrs := make([]string, len(dates))
	for i, date := range dates {
		dateStrs[i] = date.Format("2006-01-02 15:04:05")
	}
	datesParam := "'" + strings.Join(dateStrs, "','") + "'"

	query := `
		SELECT EXISTS (
			SELECT id 
			FROM appointments 
			WHERE nursing_id IN (%s)
			AND est_date IN (%s)
		)
	`
	query = fmt.Sprintf(query, nursingIdsParam, datesParam)

	var exists bool
	if err := repo.db.GetContext(ctx, &exists, query); err != nil {
		return fmt.Errorf("failed to check nurse availability: %w", err)
	}

	if exists {
		return common.ErrNursesNotAvailable
	}

	return nil
}

package categoryrepository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *categoryRepo) AddStaffToCategory(ctx context.Context, cateId, staffId uuid.UUID) error {
	query := "UPDATE categories SET staff_id = ? WHERE id=?"

	result, err := repo.db.ExecContext(ctx, query, staffId, cateId)
	if err != nil {
		return err
	}

	rowAffect, _ := result.RowsAffected()
	if rowAffect == 0 {
		return fmt.Errorf("category-id-%v doesn't exist - add staff failed", cateId)
	}
	return nil
}

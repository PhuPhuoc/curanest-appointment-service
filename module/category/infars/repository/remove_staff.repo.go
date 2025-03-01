package categoryrepository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *categoryRepo) RemoveStaffOfCategory(ctx context.Context, cateId uuid.UUID) error {
	query := "UPDATE categories SET staff_id = ? WHERE id=?"

	result, err := repo.db.ExecContext(ctx, query, nil, cateId)
	if err != nil {
		return err
	}

	rowAffect, _ := result.RowsAffected()
	if rowAffect == 0 {
		return fmt.Errorf("category-id-%v doesn't exist - remove staff failed", cateId)
	}
	return nil
}

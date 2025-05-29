package servicerepository

import (
	"context"

	"github.com/google/uuid"
)

func (repo *serviceRepo) GetCountTotalService(ctx context.Context, categoryId *uuid.UUID) (int, error) {
	var count int
	var query string
	var args []interface{}

	query = `SELECT COUNT(*) FROM services WHERE deleted_at IS NULL`

	if categoryId != nil {
		query += ` AND category_id = ?`
		args = append(args, categoryId.String())
	}

	if err := repo.db.GetContext(ctx, &count, query, args...); err != nil {
		return 0, err
	}

	return count, nil
}

package svcpackagerepository

import (
	"context"

	"github.com/google/uuid"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

func (repo *svcPackageRepo) GetSvcPackageUsageCount(ctx context.Context, categoryId uuid.UUID) ([]svcpackagedomain.ServicePackageUsage, error) {
	query := `
		SELECT 
			sp.id, sp.name,
			COUNT(cp.id) AS usage_count
		FROM 
			service_packages sp
		LEFT JOIN 
			customized_packages cp ON sp.id = cp.service_package_id
		where sp.service_id in (select s.id from services s where s.category_id = ?)
		GROUP BY 
			sp.id

		ORDER BY 
			usage_count DESC
	`

	var dtos []SvcPackageUsageCountDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, categoryId); err != nil {
		return nil, err
	}

	result := make([]svcpackagedomain.ServicePackageUsage, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToSvcPackageUsageCountEntity()
		result[i] = *entity
	}

	return result, nil
}

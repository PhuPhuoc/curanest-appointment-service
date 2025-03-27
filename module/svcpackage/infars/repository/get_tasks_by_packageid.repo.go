package svcpackagerepository

import (
	"context"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
)

func (repo *svcPackageRepo) GetServiceTasksByPackageId(ctx context.Context, svcPackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error) {
	return []svcpackagedomain.ServiceTask{}, nil
}

package svcpackagerepository

import (
	"context"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
	"github.com/google/uuid"
)

func (repo *svcPackageRepo) GetServicePackageById(ctx context.Context, svcPackageId uuid.UUID) (*svcpackagedomain.ServicePackage, error) {
	return nil, nil
}

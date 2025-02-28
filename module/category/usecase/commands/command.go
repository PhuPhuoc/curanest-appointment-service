package categorycommands

import (
	"context"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

type Commands struct {
	CreateCategory *createCategoryHandler
}

type Builder interface {
	BuildCategoryCmdRepo() CategoryCommandRepo
	BuildExternalAccountServiceInCmd() ExternalAccountService
}

func NewCategoryCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateCategory: NewCreateCategoryHandler(
			b.BuildCategoryCmdRepo(),
		),
	}
}

type CategoryCommandRepo interface {
	Create(ctx context.Context, entity *categorydomain.Category) error
	// Update(ctx context.Context, entity *categorydomain.Category) error
}

type ExternalAccountService interface {
	// UpdateAccountRoleToStaffRPC(ctx context.Context, staffId uuid.UUID) error
}

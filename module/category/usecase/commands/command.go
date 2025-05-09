package categorycommands

import (
	"context"

	"github.com/google/uuid"

	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

type Commands struct {
	CreateCategory *createCategoryHandler
	UpdateCategory *updateCategoryHandler

	AddStaff    *addStaffToCategoryHandler
	RemoveStaff *removeStaffToCategoryHandler
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
		UpdateCategory: NewUpdateCategoryHandler(
			b.BuildCategoryCmdRepo(),
		),
		AddStaff: NewAddStaffToCategoryHandler(
			b.BuildCategoryCmdRepo(),
			b.BuildExternalAccountServiceInCmd(),
		),
		RemoveStaff: NewRemoveStaffToCategoryHandler(
			b.BuildCategoryCmdRepo(),
			b.BuildExternalAccountServiceInCmd(),
		),
	}
}

type CategoryCommandRepo interface {
	Create(ctx context.Context, entity *categorydomain.Category) error
	Update(ctx context.Context, entity *categorydomain.Category) error

	AddStaffToCategory(ctx context.Context, cateId, staffId uuid.UUID) error
	RemoveStaffOfCategory(ctx context.Context, cateId uuid.UUID) error
}

type ExternalAccountService interface {
	UpdateAccountRoleRPC(ctx context.Context, nursingId uuid.UUID, targetRole string) error
}

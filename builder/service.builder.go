package builder

import (
	"github.com/jmoiron/sqlx"

	categoryrepository "github.com/PhuPhuoc/curanest-appointment-service/module/category/infars/repository"
	servicerepository "github.com/PhuPhuoc/curanest-appointment-service/module/service/infars/repository"
	servicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/commands"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
)

type builderOfService struct {
	db *sqlx.DB
}

func NewServiceBuilder(db *sqlx.DB) builderOfService {
	return builderOfService{db: db}
}

func (s builderOfService) BuildServiceCmdRepo() servicecommands.ServiceCommandRepo {
	return servicerepository.NewServiceRepo(s.db)
}

func (s builderOfService) BuildServiceQueryRepo() servicequeries.ServiceQueryRepo {
	return servicerepository.NewServiceRepo(s.db)
}

func (s builderOfService) BuildCategoryFetcher() servicequeries.CategoryFetcher {
	return categoryrepository.NewCategoryRepo(s.db)
}

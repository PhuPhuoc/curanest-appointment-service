package builder

import (
	"github.com/jmoiron/sqlx"

	categoryauthrpc "github.com/PhuPhuoc/curanest-appointment-service/module/category/infars/externalrpc/auth"
	categorynursingrpc "github.com/PhuPhuoc/curanest-appointment-service/module/category/infars/externalrpc/nursing"
	categoryrepository "github.com/PhuPhuoc/curanest-appointment-service/module/category/infars/repository"
	categorycommands "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/commands"
	categoryqueries "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/queries"
)

type builderOfCategory struct {
	db                    *sqlx.DB
	urlPathAccountService string
	urlPathNursingService string
}

func NewCategoryBuilder(db *sqlx.DB) builderOfCategory {
	return builderOfCategory{db: db}
}

func (s builderOfCategory) AddUrlPathAccountService(url string) builderOfCategory {
	s.urlPathAccountService = url
	return s
}

func (s builderOfCategory) AddUrlPathNursingService(url string) builderOfCategory {
	s.urlPathNursingService = url
	return s
}

func (s builderOfCategory) BuildCategoryCmdRepo() categorycommands.CategoryCommandRepo {
	return categoryrepository.NewCategoryRepo(s.db)
}

func (s builderOfCategory) BuildExternalAccountServiceInCmd() categorycommands.ExternalAccountService {
	return categoryauthrpc.NewAccountRPC(s.urlPathAccountService)
}

func (s builderOfCategory) BuildCategoryQueryRepo() categoryqueries.CategoryQueryRepo {
	return categoryrepository.NewCategoryRepo(s.db)
}

func (s builderOfCategory) BuildExternalNursingServiceInQuery() categoryqueries.ExternalNursingService {
	return categorynursingrpc.NewNursingRPC(s.urlPathNursingService)
}

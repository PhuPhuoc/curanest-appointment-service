package builder

import (
	"github.com/jmoiron/sqlx"

	cuspackagerepository "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/infars/repository"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
	cuspackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/queries"
)

type builderOfCusPackage struct {
	db *sqlx.DB
}

func NewCusPackageBuilder(db *sqlx.DB) builderOfCusPackage {
	return builderOfCusPackage{db: db}
}

func (s builderOfCusPackage) BuildCusPackageCmdRepo() cuspackagecommands.CusPackageCommandRepo {
	return cuspackagerepository.NewCusPackageRepo(s.db)
}

func (s builderOfCusPackage) BuildCusPackageQueryRepo() cuspackagequeries.CusPackageQueryRepo {
	return cuspackagerepository.NewCusPackageRepo(s.db)
}

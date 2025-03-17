package builder

import (
	"github.com/jmoiron/sqlx"

	svcpackagerepository "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/infars/repository"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
	svcpackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/queries"
)

type builderOfSvcPackage struct {
	db *sqlx.DB
}

func NewSvcPackageBuilder(db *sqlx.DB) builderOfSvcPackage {
	return builderOfSvcPackage{db: db}
}

func (s builderOfSvcPackage) BuildSvcPackageCmdRepo() svcpackagecommands.SvcPackageCommandRepo {
	return svcpackagerepository.NewSvcPackageRepo(s.db)
}

func (s builderOfSvcPackage) BuildSvcPackageQueryRepo() svcpackagequeries.SvcPackageQueryRepo {
	return svcpackagerepository.NewSvcPackageRepo(s.db)
}

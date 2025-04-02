package builder

import (
	"github.com/jmoiron/sqlx"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentrepository "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/infars/repository"
	cuspackagerepository "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/infars/repository"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
	cuspackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/queries"
	invoicerepository "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/infars/repository"
	svcpackagerepository "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/infars/repository"
)

type builderOfCusPackage struct {
	db             *sqlx.DB
	transactionMgr common.TransactionManager
}

func NewCusPackageBuilder(db *sqlx.DB) builderOfCusPackage {
	return builderOfCusPackage{
		db:             db,
		transactionMgr: NewSQLxTransactionManager(db),
	}
}

func (s builderOfCusPackage) BuildCusPackageCmdRepo() cuspackagecommands.CusPackageCommandRepo {
	return cuspackagerepository.NewCusPackageRepo(s.db)
}

func (s builderOfCusPackage) BuildCusPackageQueryRepo() cuspackagequeries.CusPackageQueryRepo {
	return cuspackagerepository.NewCusPackageRepo(s.db)
}

func (s builderOfCusPackage) BuildSvcPackageFetcher() cuspackagecommands.SvcPackageFetcher {
	return svcpackagerepository.NewSvcPackageRepo(s.db)
}

func (s builderOfCusPackage) BuildAppointmentFetcher() cuspackagecommands.AppointmentFetcher {
	return appointmentrepository.NewAppointmentRepo(s.db)
}

func (s builderOfCusPackage) BuildInvoiceFetcher() cuspackagecommands.InvoiceFetcher {
	return invoicerepository.NewInvoiceRepo(s.db)
}

func (s builderOfCusPackage) BuildTransactionManager() common.TransactionManager {
	return s.transactionMgr
}

package builder

import (
	"github.com/jmoiron/sqlx"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentrepository "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/infars/repository"
	externalapigoong "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/infars/externalapi"
	cuspackagerepository "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/infars/repository"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
	cuspackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/queries"
	invoicerepository "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/infars/repository"
	svcpackagerepository "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/infars/repository"
)

type builderOfCusPackage struct {
	db             *sqlx.DB
	transactionMgr common.TransactionManager
	payOS          common.PayOSConfig

	goongApiUrl string
	goongApiKey string
}

func NewCusPackageBuilder(db *sqlx.DB) builderOfCusPackage {
	return builderOfCusPackage{
		db:             db,
		transactionMgr: NewSQLxTransactionManager(db),
	}
}

func (s builderOfCusPackage) AddPayOsConfig(payOS common.PayOSConfig) builderOfCusPackage {
	s.payOS = payOS
	return s
}

func (s builderOfCusPackage) AddGoongConfig(apiurl, apikey string) builderOfCusPackage {
	s.goongApiUrl = apiurl
	s.goongApiKey = apikey
	return s
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

func (s builderOfCusPackage) BuilderPayosConfig() common.PayOSConfig {
	return s.payOS
}

func (s builderOfCusPackage) BuildExternalGoongAPI() cuspackagecommands.ExternalGoongAPI {
	return externalapigoong.NewExternalGoongAPI(s.goongApiUrl, s.goongApiKey)
}

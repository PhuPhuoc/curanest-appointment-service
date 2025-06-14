package builder

import (
	"github.com/jmoiron/sqlx"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	externalapi "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/infars/expertnalapi"
	appointmentrepository "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/infars/repository"
	apppointmentcommands "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/commands"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
	cuspackagerepository "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/infars/repository"
	servicerepository "github.com/PhuPhuoc/curanest-appointment-service/module/service/infars/repository"
)

type builderOfAppointment struct {
	db                 *sqlx.DB
	urlPushNotiService string
	urlNurseService    string
	urlPatientService  string

	transactionMgr common.TransactionManager
	// goongApiUrl string
	// goongApiKey string
}

func (s builderOfAppointment) AddPathPushNotiService(url string) builderOfAppointment {
	s.urlPushNotiService = url
	return s
}

func (s builderOfAppointment) AddPathPatientService(url string) builderOfAppointment {
	s.urlPatientService = url
	return s
}

func (s builderOfAppointment) AddPathUrlNursingService(url string) builderOfAppointment {
	s.urlNurseService = url
	return s
}

// func (s builderOfAppointment) AddGoongConfig(apiurl, apikey string) builderOfAppointment {
// 	s.goongApiUrl = apiurl
// 	s.goongApiKey = apikey
// 	return s
// }

func NewAppointmentBuilder(db *sqlx.DB) builderOfAppointment {
	return builderOfAppointment{
		db:             db,
		transactionMgr: NewSQLxTransactionManager(db),
	}
}

func (s builderOfAppointment) BuildTransactionManager() common.TransactionManager {
	return s.transactionMgr
}

func (s builderOfAppointment) BuildAppointmentCmdRepo() apppointmentcommands.AppointmentCommandRepo {
	return appointmentrepository.NewAppointmentRepo(s.db)
}

func (s builderOfAppointment) BuilderCusPackageFetcher() appointmentqueries.CusPackageFetcher {
	return cuspackagerepository.NewCusPackageRepo(s.db)
}

func (s builderOfAppointment) BuilderServiceFetcher() appointmentqueries.ServiceFetcher {
	return servicerepository.NewServiceRepo(s.db)
}

func (s builderOfAppointment) BuildAppointmentQueryRepo() appointmentqueries.AppointmentQueryRepo {
	return appointmentrepository.NewAppointmentRepo(s.db)
}

func (s builderOfAppointment) BuildCusTaskFetcher() apppointmentcommands.CusTaskFetcher {
	return cuspackagerepository.NewCusPackageRepo(s.db)
}

func (s builderOfAppointment) BuildMedicalRecord() apppointmentcommands.MedicalRecordFetcher {
	return cuspackagerepository.NewCusPackageRepo(s.db)
}

func (s builderOfAppointment) BuildExternalNurseServiceApiCmd() apppointmentcommands.ExternalNursingService {
	return externalapi.NewNursingRPC(s.urlNurseService)
}

func (s builderOfAppointment) BuildExternalNurseServiceApiQuery() appointmentqueries.ExternalNursingService {
	return externalapi.NewNursingRPC(s.urlNurseService)
}

// func (s builderOfAppointment) BuildExternalGoongAPI() apppointmentcommands.ExternalGoongAPI {
// 	return externalapi.NewExternalGoongAPI(s.goongApiUrl, s.goongApiKey)
// }

func (s builderOfAppointment) BuildExternalPushNotiService() apppointmentcommands.ExternalPushNotiService {
	return externalapi.NewPushNotiServiceRPC(s.urlPushNotiService)
}

func (s builderOfAppointment) BuildExternalPatientService() apppointmentcommands.ExternalPatientService {
	return externalapi.NewPatientServiceAPI(s.urlPatientService)
}

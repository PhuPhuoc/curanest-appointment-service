package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/PhuPhuoc/curanest-appointment-service/builder"
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/PhuPhuoc/curanest-appointment-service/config"
	"github.com/PhuPhuoc/curanest-appointment-service/docs"
	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	appointmenthttpservice "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/infars/httpservice"
	apppointmentcommands "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/commands"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
	categoryhttpservice "github.com/PhuPhuoc/curanest-appointment-service/module/category/infars/httpservice"
	categorycommands "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/commands"
	categoryqueries "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/queries"
	cuspackagehttpservice "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/infars/httpservice"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
	cuspackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/queries"
	invoicehttpservice "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/infars/httpservice"
	invoicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/commands"
	invoicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/queries"
	servicehttpservice "github.com/PhuPhuoc/curanest-appointment-service/module/service/infars/httpservice"
	servicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/commands"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
	svcpackagehttpservice "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/infars/httpservice"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
	svcpackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/queries"
)

type server struct {
	port string
	db   *sqlx.DB
}

func InitServer(port string, db *sqlx.DB) *server {
	return &server{
		port: port,
		db:   db,
	}
}

const (
	env_local     = "local"
	env_vps       = "vps"
	url_acc_local = "http://localhost:8001"
	url_acc_prod  = "http://auth_service:8080"

	url_nursing_local = "http://localhost:8003"
	url_nursing_prod  = "http://nurse_service:8080"
)

// @Summary		ping server
// @Description	ping server
// @Tags			ping
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]any	"message success"
// @Failure		400	{object}	error			"Bad request error"
// @Router			/ping [get]
func (sv *server) RunApp() error {
	var urlAccServices string
	var urlNursingServices string
	envDevlopment := config.AppConfig.EnvDev
	goongApiUrl := config.AppConfig.GoongAPIURL
	goongApiKey := config.AppConfig.GoongAPIKEY
	if envDevlopment == env_local {
		// gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/"
		urlAccServices = url_acc_local
		urlNursingServices = url_nursing_local
	}

	if envDevlopment == env_vps {
		gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/appointment"
		urlAccServices = url_acc_prod
		urlNursingServices = url_nursing_prod
	}

	router := gin.New()

	configcors := cors.DefaultConfig()
	configcors.AllowAllOrigins = true
	configcors.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "PATCH", "OPTIONS"}
	configcors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	configcors.ExposeHeaders = []string{"Content-Length"}
	configcors.AllowCredentials = true
	configcors.MaxAge = 12 * time.Hour

	router.Use(cors.New(configcors))
	router.Use(middleware.SkipSwaggerLog(), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "curanest-appointment-service - pong"}) })

	authClient := common.NewJWTx(config.AppConfig.Key)
	payosConfig := common.NewPayOs(config.AppConfig.PayOsClientId, config.AppConfig.PayOsApiKey, config.AppConfig.PayOsCheckSumKey)

	category_cmd_builder := categorycommands.NewCategoryCmdWithBuilder(
		builder.NewCategoryBuilder(sv.db).AddUrlPathAccountService(urlAccServices),
	)
	category_query_builder := categoryqueries.NewCategoryQueryWithBuilder(
		builder.NewCategoryBuilder(sv.db).AddUrlPathNursingService(urlNursingServices),
	)

	service_cmd_builder := servicecommands.NewServiceCmdWithBuilder(
		builder.NewServiceBuilder(sv.db),
	)
	service_query_builder := servicequeries.NewServiceQueryWithBuilder(
		builder.NewServiceBuilder(sv.db),
	)

	svcpackage_cmd_builder := svcpackagecommands.NewSvcPackageCmdWithBuilder(
		builder.NewSvcPackageBuilder(sv.db),
	)
	svcpackage_query_builder := svcpackagequeries.NewSvcPackageQueryWithBuilder(
		builder.NewSvcPackageBuilder(sv.db),
	)

	cuspackage_cmd_builder := cuspackagecommands.NewCusPackageCmdWithBuilder(
		builder.NewCusPackageBuilder(sv.db).AddPayOsConfig(*payosConfig).AddGoongConfig(goongApiUrl, goongApiKey),
	)
	cuspackage_query_builder := cuspackagequeries.NewCusPackageQueryWithBuilder(
		builder.NewCusPackageBuilder(sv.db),
	)

	appointment_cmd_builder := apppointmentcommands.NewAppointmentCmdWithBuilder(
		builder.NewAppointmentBuilder(sv.db),
	)
	appointment_query_builder := appointmentqueries.NewAppointmentQueryWithBuilder(

		builder.NewAppointmentBuilder(sv.db),
	)

	invoice_cmd_builder := invoicecommands.NewInvoiceCmdWithBuilder(
		builder.NewInvoiceBuilder(sv.db).AddPayOsConfig(*payosConfig),
	)
	invoice_query_builder := invoicequeries.NewInvoiceQueryWithBuilder(
		builder.NewInvoiceBuilder(sv.db),
	)

	api := router.Group("/api/v1")
	{
		categoryhttpservice.NewCategoryHTTPService(
			category_cmd_builder,
			category_query_builder,
		).AddAuth(authClient).Routes(api)

		servicehttpservice.NewServiceHTTPService(
			service_cmd_builder,
			service_query_builder,
		).AddAuth(authClient).Routes(api)

		svcpackagehttpservice.NewSvcPackageHTTPService(
			svcpackage_cmd_builder,
			svcpackage_query_builder,
		).AddAuth(authClient).Routes(api)

		cuspackagehttpservice.NewSvcPackageHTTPService(
			cuspackage_cmd_builder,
			cuspackage_query_builder,
		).AddAuth(authClient).Routes(api)

		appointmenthttpservice.NewAppointmentHTTPService(
			appointment_cmd_builder,
			appointment_query_builder,
		).AddAuth(authClient).Routes(api)

		invoicehttpservice.NewInvoiceHTTPService(
			invoice_cmd_builder,
			invoice_query_builder,
		).AddAuth(authClient).Routes(api)
	}

	// rpc := router.Group("/external/rpc")
	// {
	// }

	log.Println("server start listening at port: ", sv.port)
	return router.Run(sv.port)
}

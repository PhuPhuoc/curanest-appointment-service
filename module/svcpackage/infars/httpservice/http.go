package svcpackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	svcpackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/commands"
	svcpackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/usecase/queries"
	"github.com/gin-gonic/gin"
)

type svcPackageHttpService struct {
	command svcpackagecommands.Commands
	query   svcpackagequeries.Queries
	auth    middleware.AuthClient
}

func NewSvcPackageHTTPService(command svcpackagecommands.Commands, query svcpackagequeries.Queries) *svcPackageHttpService {
	return &svcPackageHttpService{
		command: command,
		query:   query,
	}
}

func (s *svcPackageHttpService) AddAuth(auth middleware.AuthClient) *svcPackageHttpService {
	s.auth = auth
	return s
}

func (s *svcPackageHttpService) Routes(g *gin.RouterGroup) {
	service_svcpackage_route := g.Group("/services/:service-id")
	{
		service_svcpackage_route.POST(
			"/svcpackage",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("staff"),
			s.handleCreateServicePackage(),
		)
		service_svcpackage_route.PUT(
			"/svcpackage/:svcpackage-id",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("staff"),
			s.handleUpdateServicePackage(),
		)
		service_svcpackage_route.GET(
			"/svcpackage",
			s.handleGetServicPackage(),
		)
	}

	svcpackage_route := g.Group("/svcpackage")
	{

		svcpackage_route.POST(
			"/:svcpackage-id/svctask",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("staff"),
			s.handleCreateServiceTask(),
		)
		svcpackage_route.PUT(
			"/:svcpackage-id/svctask/:svctask-id",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("staff"),
			s.handleUpdateServiceTask(),
		)
		svcpackage_route.GET(
			"/:svcpackage-id/svctask",
			s.handleGetServicTasks(),
		)
	}
}

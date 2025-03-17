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
	svcpackage_route := g.Group("/svcpackage")
	{
		svcpackage_route.POST(
			"",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("staff"),
			s.handleCreateServicePackage(),
		)
		svcpackage_route.POST(
			"/:svcpackage-id/svctask",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("staff"),
			s.handleCreateServiceTask(),
		)
	}
}

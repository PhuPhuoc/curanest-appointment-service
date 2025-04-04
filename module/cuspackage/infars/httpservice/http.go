package cuspackagehttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
	cuspackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/queries"
)

type cusPackageHttpService struct {
	cmd   cuspackagecommands.Commands
	query cuspackagequeries.Queries
	auth  middleware.AuthClient
}

func NewSvcPackageHTTPService(cmd cuspackagecommands.Commands, query cuspackagequeries.Queries) *cusPackageHttpService {
	return &cusPackageHttpService{
		cmd:   cmd,
		query: query,
	}
}

func (s *cusPackageHttpService) AddAuth(auth middleware.AuthClient) *cusPackageHttpService {
	s.auth = auth
	return s
}

func (s *cusPackageHttpService) Routes(g *gin.RouterGroup) {
	cuspackage_route := g.Group("/cuspackage")
	{
		cuspackage_route.POST(
			"",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("relatives"),
			s.handleCreateCustomizedPackageAndTask(),
		)
	}
}

package servicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	servicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/commands"
	servicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/queries"
	"github.com/gin-gonic/gin"
)

type serviceHttpService struct {
	command servicecommands.Commands
	query   servicequeries.Queries
	auth    middleware.AuthClient
}

func NewServiceHTTPService(command servicecommands.Commands, query servicequeries.Queries) *serviceHttpService {
	return &serviceHttpService{
		command: command,
		query:   query,
	}
}

func (s *serviceHttpService) AddAuth(auth middleware.AuthClient) *serviceHttpService {
	s.auth = auth
	return s
}

func (s *serviceHttpService) Routes(g *gin.RouterGroup) {
	category_service_route := g.Group("/categories/:category-id")
	{
		category_service_route.GET(
			"/services",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleGetServiceByCategory(),
		)
	}

	service_route := g.Group("/services")
	{
		service_route.POST(
			"",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleCreateService(),
		)
	}
}

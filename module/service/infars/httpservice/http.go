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
		category_service_route.POST(
			"/services",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleCreateService(),
		)
	}

	staff_service_route := g.Group("/staff/services")
	{
		staff_service_route.GET(
			"",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("staff"),
			s.handleGetServiceOfStaff(),
		)
	}

	service_route := g.Group("/services")
	{
		service_route.GET(
			"/group-by-category",
			s.handleGetServiceGroupByCategory(),
		)
		service_route.GET(
			"/staff",
			s.handleGetServiceOfStaff(),
		)
		service_route.PUT(
			"",
			s.handleUpdateService(),
		)
	}
}

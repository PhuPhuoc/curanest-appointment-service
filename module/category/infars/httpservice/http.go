package categoryhttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	categorycommands "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/commands"
	categoryqueries "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/queries"
)

type categoryHttpService struct {
	command categorycommands.Commands
	query   categoryqueries.Queries
	auth    middleware.AuthClient
}

func NewCategoryHTTPService(command categorycommands.Commands, query categoryqueries.Queries) *categoryHttpService {
	return &categoryHttpService{
		command: command,
		query:   query,
	}
}

func (s *categoryHttpService) AddAuth(auth middleware.AuthClient) *categoryHttpService {
	s.auth = auth
	return s
}

func (s *categoryHttpService) Routes(g *gin.RouterGroup) {
	cate_route := g.Group("/categories")
	{
		cate_route.POST(
			"",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleCreateCategory(),
		)
		cate_route.GET(
			"",
			middleware.RequireAuth(s.auth),
			s.handleGetCategories(),
		)
	}
}

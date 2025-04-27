package appointmenthttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	apppointmentcommands "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/commands"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

type appointmentHttpService struct {
	command apppointmentcommands.Commands
	query   appointmentqueries.Queries
	auth    middleware.AuthClient
}

func NewAppointmentHTTPService(command apppointmentcommands.Commands, query appointmentqueries.Queries) *appointmentHttpService {
	return &appointmentHttpService{
		command: command,
		query:   query,
	}
}

func (s *appointmentHttpService) AddAuth(auth middleware.AuthClient) *appointmentHttpService {
	s.auth = auth
	return s
}

func (s *appointmentHttpService) Routes(g *gin.RouterGroup) {
	appointment_route := g.Group("/appointments")
	{
		appointment_route.GET(
			"",
			// middleware.RequireAuth(s.auth),
			s.handleGetAppointmentByFilter(),
		)
		appointment_route.GET(
			"/nursing-timesheet",
			// middleware.RequireAuth(s.auth),
			s.handleGetNursingTimesheet(),
		)
		appointment_route.GET(
			"/nursing-available",
			// middleware.RequireAuth(s.auth),
			s.handleGetNursingAvailable(),
		)
		appointment_route.PATCH(
			":appointment-id/status",
			// middleware.RequireAuth(s.auth),
			s.handleChangeAppointmentStatus(),
		)
		appointment_route.POST(
			"/verify-nurses-dates",
			s.handleVerifyNurseWithDate(),
		)
	}
}

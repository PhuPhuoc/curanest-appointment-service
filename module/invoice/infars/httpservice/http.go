package invoicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	invoicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/commands"
	invoicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/queries"
	"github.com/gin-gonic/gin"
)

type invoiceHttpService struct {
	cmd   invoicecommands.Commands
	query invoicequeries.Queries
	auth  middleware.AuthClient
}

func NewInvoiceHTTPService(cmd invoicecommands.Commands, query invoicequeries.Queries) *invoiceHttpService {
	return &invoiceHttpService{
		cmd:   cmd,
		query: query,
	}
}

func (s *invoiceHttpService) AddAuth(auth middleware.AuthClient) *invoiceHttpService {
	s.auth = auth
	return s
}

func (s *invoiceHttpService) Routes(g *gin.RouterGroup) {
	cuspackage_invoice_route := g.Group("/cuspackage/:cus-package-id/invoices")
	{
		cuspackage_invoice_route.GET(
			"",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("relatives"),
			s.handleFindInvoice(),
		)
	}
	// invoice_route := g.Group("/invoices")
	// {
	// 	invoice_route.GET(
	// 		"/:invoice-id/url-payment",
	// 		// middleware.RequireAuth(s.auth),
	// 		// middleware.RequireRole("relatives"),
	// 		s.handleGetUrlPaymentForInvoice(),
	// 	)
	// }
}

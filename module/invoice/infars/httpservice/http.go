package invoicehttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	invoicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/commands"
	invoicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/queries"
)

type invoiceHttpService struct {
	cmd   invoicecommands.Commands
	query invoicequeries.Queries
	auth  middleware.AuthClient

	checksumKey string
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

func (s *invoiceHttpService) AddChecksumKey(checksumKey string) *invoiceHttpService {
	s.checksumKey = checksumKey
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
	invoice_route := g.Group("/invoices")
	{

		invoice_route.POST(
			"/revenue",
			s.handleGetRevenue(),
		)
		invoice_route.POST(
			"/webhooks",
			s.handlePayosWebhook(),
		)
		invoice_route.POST("/patient",
			s.handleInvoicesByPatientIds(),
		)
		invoice_route.PATCH(
			"/cancel-payment-url/:order-code",
			s.handleCancelPaymentUrl(),
		)
		invoice_route.PATCH(
			"/:invoice-id/create-payment-url",
			s.handleCreateNewPaymentUrl(),
		)
	}
}

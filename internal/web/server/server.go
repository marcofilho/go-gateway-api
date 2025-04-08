package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marcofilho/go-api-payment-gateway/internal/service"
	"github.com/marcofilho/go-api-payment-gateway/internal/web/handlers"
	"github.com/marcofilho/go-api-payment-gateway/internal/web/middleware"
)

type Server struct {
	httpServer     *http.Server
	router         *chi.Mux
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	router := chi.NewRouter()

	return &Server{
		router:         router,
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) SetupRoutes() {
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService, s.accountService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		s.router.Post("/invoice", invoiceHandler.Create)
		s.router.Get("/invoice/{id}", invoiceHandler.GetByID)
		s.router.Get("/invoice", invoiceHandler.GetInvoicesByAccountID)
	})

}

func (s *Server) Start() error {
	s.httpServer = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.httpServer.ListenAndServe()
}

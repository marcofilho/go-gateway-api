package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marcofilho/go-api-payment-gateway/internal/service"
	"github.com/marcofilho/go-api-payment-gateway/internal/web/handlers"
)

type Server struct {
	httpServer     *http.Server
	router         *chi.Mux
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	router := chi.NewRouter()

	return &Server{
		router:         router,
		accountService: accountService,
		port:           port,
	}
}

func (s *Server) SetupRoutes() {
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

}

func (s *Server) Start() error {
	s.httpServer = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.httpServer.ListenAndServe()
}

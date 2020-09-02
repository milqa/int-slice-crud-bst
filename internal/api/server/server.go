package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/milQA/int-slice-crud-bst/internal/api/service"
	"go.uber.org/zap"
)

type (
	Server struct {
		router  *chi.Mux
		service *service.Service
		logger  *zap.Logger
	}
)

func NewServer(log *zap.Logger, service *service.Service) *Server {

	router := chi.NewMux()

	logger := log.With(
		zap.String(
			"module", "http",
		),
	)

	server := &Server{
		router:  router,
		service: service,
		logger:  logger,
	}

	router.Use(LogRequestMiddleware(logger))
	router.MethodFunc(http.MethodPost, "/insert", server.insert())
	router.MethodFunc(http.MethodDelete, "/delete", server.delete())
	router.MethodFunc(http.MethodGet, "/search", server.search())

	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	s.router.ServeHTTP(w, r)
}

package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/milQA/int-slice-crud-bst/internal/api/store"
	"go.uber.org/zap"
)

type (
	Server struct {
		router *chi.Mux
		store  store.Store
		logger *zap.Logger
	}
)

func NewServer(log *zap.Logger, store store.Store) *Server {

	router := chi.NewMux()

	logger := log.With(
		zap.String(
			"module", "http",
		),
	)

	router.Use(LogRequestMiddleware(logger))
	router.MethodFunc(http.MethodPost, "/insert", insert(store, logger))
	router.MethodFunc(http.MethodDelete, "/delete", delete(store, logger))
	router.MethodFunc(http.MethodGet, "/search", search(store, logger))

	server := &Server{
		router: router,
		store:  store,
		logger: logger,
	}

	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	s.router.ServeHTTP(w, r)
}

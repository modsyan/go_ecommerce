package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/go-chi/chi"
)

type APIServer struct {
	addr   string
	router *mux.Router
	// router     chi.Router
	db         *sql.DB
	routes     []RouteConfig
	basePrefix string
}

func NewAPIServer(addr string, db *sql.DB, basePrefix string) *APIServer {
	return &APIServer{
		addr:       addr,
		router:     mux.NewRouter(),
		db:         db,
		basePrefix: basePrefix,
	}
}

func (s *APIServer) Run() error {
	log.Println("Listening on", s.addr)

	s.SetupRoutes()
	s.applyMiddleware()

	return http.ListenAndServe(s.addr, s.router)
}

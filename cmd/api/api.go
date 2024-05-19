package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr   string
	router *mux.Router
	db     *sql.DB
	routes []RouteConfig
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr:   addr,
		router: mux.NewRouter(),
		db:     db,
	}
}

func (s *APIServer) Run() error {
	log.Println("Listening on", s.addr)

	s.SetupRoutes()
	s.applyMiddleware()

	return http.ListenAndServe(s.addr, s.router)
}

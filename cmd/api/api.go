package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/modsyan/go_ecommerce/modules/cart"
	"github.com/modsyan/go_ecommerce/modules/healthz"
	"github.com/modsyan/go_ecommerce/modules/product"
	"github.com/modsyan/go_ecommerce/modules/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	healthzHandler := healthz.CreateHandler()
	healthzHandler.RegisterRoutes(subrouter)

	userHandler := user.CreateHandler()
	userHandler.RegisterRoutes(subrouter)

	cartHandler := cart.CreateHandler()
	cartHandler.RegisterRoutes(subrouter)

	productHandler := product.CreateHandler()
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

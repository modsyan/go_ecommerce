package api

import (
	"github.com/gorilla/mux"
	"github.com/modsyan/go_ecommerce/internal/cart"
	"github.com/modsyan/go_ecommerce/internal/healthz"
	"github.com/modsyan/go_ecommerce/internal/product"
	"github.com/modsyan/go_ecommerce/internal/user"
)

type RouteConfig struct {
	Register func(router *mux.Router)
	Prefix   string
}

func intializeRoutes() []RouteConfig {
	return []RouteConfig{
		{Prefix: "/healthz", Register: healthz.CreateHandler().RegisterRoutes},
		{Prefix: "/user", Register: user.CreateHandler().RegisterRoutes},
		{Prefix: "/cart", Register: cart.CreateHandler().RegisterRoutes},
		{Prefix: "/product", Register: product.CreateHandler().RegisterRoutes},
	}
}

func (s *APIServer) SetupRoutes() {
	s.routes = intializeRoutes()
	for _, route := range s.routes {
		fullPath := s.basePrefix + route.Prefix
		subRouter := s.router.PathPrefix(fullPath).Subrouter()
		route.Register(subRouter)
	}
}

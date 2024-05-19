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

const basePrefixer = "/api/v1"

// func (s *APIServer) AddRoute(prefix string, register func(router *mux.Router)) {
// 	s.routes = append(s.routes, RouteConfig{Register: register, Prefix: prefix})
// }
//
// func (s *APIServer) IntializeRoutes() {
// 	s.AddRoute("/healhz", healthz.CreateHandler().RegisterRoutes)
// 	s.AddRoute("/user", user.CreateHandler().RegisterRoutes)
// 	s.AddRoute("/cart", cart.CreateHandler().RegisterRoutes)
// 	s.AddRoute("/product", product.CreateHandler().RegisterRoutes)
// }

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
		path := basePrefixer + route.Prefix
		subRouter := s.router.PathPrefix(path).Subrouter()
		route.Register(subRouter)
	}
}

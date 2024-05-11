package product

import (
	"github.com/gorilla/mux"
)

type Handler struct{}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/products")
	router.HandleFunc("GET /", h.GetAllHandler)
}

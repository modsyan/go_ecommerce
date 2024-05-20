package product

import (
	"github.com/gorilla/mux"
)

type Handler struct{}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.GetAllProductsHandler).Methods("GET")
	router.HandleFunc("/{id}", h.GetProductDetailsHandler).Methods("GET")
	router.HandleFunc("/", h.AddProductHandler).Methods("POST")
	router.HandleFunc("/{id}", h.RemoveProductHandler).Methods("DELETE")
	router.HandleFunc("/{id}", h.EditProductHandler).Methods("PUT")
}

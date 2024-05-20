package cart

import (
	"github.com/gorilla/mux"
)

type Handler struct{}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.GetCartlandler).Methods("GET")
	router.HandleFunc("/", h.AddItemCartHandler).Methods("POST")
	router.HandleFunc("/", h.EditCartHandler).Methods("PUT")
	router.HandleFunc("/", h.RemoveItemCartHandler).Methods("DELETE")
}

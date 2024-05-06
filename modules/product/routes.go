package product

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/products")

	router.HandleFunc("/", h.GetAllHandler)
}

func (h *Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
}

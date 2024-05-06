package cart

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/carts")

	router.HandleFunc("/", h.GetCartHandler)
}

func (h *Handler) GetCartHandler(w http.ResponseWriter, r *http.Request) {
}

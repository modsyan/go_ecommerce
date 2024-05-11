package healthz

import (
	"github.com/gorilla/mux"
)

type Handler struct{}

type HealthResponse struct {
	Status string `json:"status"`
}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/healthz", h.HealthzHandler)
}

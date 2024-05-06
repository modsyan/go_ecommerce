package healthz

import (
	"encoding/json"
	"net/http"

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

func (h *Handler) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	health := HealthResponse{Status: "healthy"}

	jsonResponse, err := json.Marshal(health)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

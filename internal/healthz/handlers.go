package healthz

import (
	"encoding/json"
	"net/http"
)

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

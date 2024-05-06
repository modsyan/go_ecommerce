package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.LoginHandler)

	router.HandleFunc("/register", h.RegisterHandler)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
}

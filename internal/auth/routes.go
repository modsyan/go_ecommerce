package user

import (
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

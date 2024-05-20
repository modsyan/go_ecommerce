package user

import (
	"net/http"

	"github.com/modsyan/go_ecommerce/internal/auth/dtos"
)

type RegisterRequest struct {
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
}

type RegisterResponse struct {
	User dtos.UserResponseDto `json:"user"`
}

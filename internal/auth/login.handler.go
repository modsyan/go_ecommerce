package user

import (
	"net/http"
)

type LoginRequest struct {
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
}

type LoginResponse struct {
	Token        string `json:"token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

package auth

import "github.com/gnomedevreact/CombatIntel/internal/models"

type authResponse struct {
	models.User
	AccessToken string `json:"access_token"`
}

type regsiterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

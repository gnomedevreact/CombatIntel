package middleware

import (
	"github.com/gnomedevreact/CombatIntel/internal/api"
	"net/http"
)

func AdminMiddleware(next http.Handler, cfg *api.ApiConfig) http.Handler {
	return AuthMiddleware(RolesMiddleware(next, []Role{"admin"}, cfg))
}

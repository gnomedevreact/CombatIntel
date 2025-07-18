package routes

import (
	"github.com/gnomedevreact/CombatIntel/internal/api"
	"github.com/gnomedevreact/CombatIntel/internal/middleware"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux, apiCfg *api.ApiConfig) {
	staticHandler := http.StripPrefix("/app/", http.FileServer(http.Dir("./static/")))

	mux.Handle("/app/", staticHandler)

	//Users
	mux.Handle("GET /users", middleware.AuthMiddleware(middleware.RolesMiddleware(http.HandlerFunc(apiCfg.GetAllUsers), []middleware.Role{"admin"}, apiCfg)))

	//Auth
	mux.HandleFunc("POST /auth/register", apiCfg.Register)
	mux.HandleFunc("POST /auth/login", apiCfg.Login)
}

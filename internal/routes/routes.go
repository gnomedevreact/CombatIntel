package routes

import (
	"github.com/gnomedevreact/CombatIntel/internal/api"
	"github.com/gnomedevreact/CombatIntel/internal/api/auth"
	"github.com/gnomedevreact/CombatIntel/internal/api/units"
	"github.com/gnomedevreact/CombatIntel/internal/api/users"
	"github.com/gnomedevreact/CombatIntel/internal/middleware"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux, apiCfg *api.ApiConfig) {
	validator := validator.New()

	staticHandler := http.StripPrefix("/app/", http.FileServer(http.Dir("./static/")))

	mux.Handle("/app/", staticHandler)

	//Users
	usersHandler := users.NewHandler(apiCfg.Db, validator)
	mux.Handle("GET /users", middleware.AdminMiddleware(http.HandlerFunc(usersHandler.GetAllUsers), apiCfg))

	//Auth
	authHandler := auth.NewHandler(apiCfg, validator)
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	//Units
	unitsHandler := units.NewHandler(apiCfg.Db, validator)
	mux.Handle("POST /units", Chain(http.HandlerFunc(unitsHandler.CreateUnit), middleware.AuthMiddleware))
}

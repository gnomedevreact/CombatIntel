package auth

import (
	"encoding/json"
	"github.com/gnomedevreact/CombatIntel/internal/api"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	authService authService
	validator   *validator.Validate
}

func NewHandler(apiCfg *api.ApiConfig, validator *validator.Validate) *Handler {
	return &Handler{
		authService: authService{cfg: apiCfg},
		validator:   validator,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var authData regsiterRequest
	err := decoder.Decode(&authData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.validator.Struct(authData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.authService.register(r.Context(), authData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var authData loginRequest
	err := decoder.Decode(&authData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.validator.Struct(authData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.authService.login(r.Context(), authData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

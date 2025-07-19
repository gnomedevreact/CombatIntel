package users

import (
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	usersService usersService
	validator    *validator.Validate
}

func NewHandler(db *database.Queries, validator *validator.Validate) *Handler {
	return &Handler{
		usersService: usersService{db},
		validator:    validator,
	}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usersService.GetAllUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}

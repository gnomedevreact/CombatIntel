package api

import (
	"github.com/gnomedevreact/CombatIntel/internal/models"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/jinzhu/copier"
	"net/http"
)

func (cfg *ApiConfig) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var usersResp []models.User
	users, err := cfg.Db.GetUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	copier.Copy(&usersResp, &users)
	utils.RespondWithJSON(w, http.StatusOK, usersResp)
}

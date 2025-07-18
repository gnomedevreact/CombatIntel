package api

import (
	"encoding/json"
	"errors"
	"github.com/gnomedevreact/CombatIntel/internal/auth"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/gnomedevreact/CombatIntel/internal/models"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type authResponse struct {
	models.User
	AccessToken string `json:"access_token"`
}

var validate = validator.New()

func (cfg *ApiConfig) Register(w http.ResponseWriter, r *http.Request) {
	type RegsiterRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,min=8"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var authData RegsiterRequest
	err := decoder.Decode(&authData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if err := validate.Struct(authData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	hashesPassword, err := bcrypt.GenerateFromPassword([]byte(authData.Password), 15)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	var userResp models.User
	user, err := cfg.Db.CreateUser(r.Context(), database.CreateUserParams{
		Username:       authData.Username,
		Password:       string(hashesPassword),
		ClearanceLevel: "C",
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	copier.Copy(&userResp, &user)

	token, err := auth.GenerateJWT(user.ID.String(), cfg.Secret)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, authResponse{
		User:        userResp,
		AccessToken: token,
	})
}

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var authData LoginRequest
	err := decoder.Decode(&authData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if err := validate.Struct(authData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	var userResp models.User
	user, err := cfg.Db.GetUserByUsername(r.Context(), authData.Username)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	copier.Copy(&userResp, &user)

	hashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(authData.Password))
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, errors.New("password incorrect"))
		return
	}

	token, err := auth.GenerateJWT(user.ID.String(), cfg.Secret)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, authResponse{
		User:        userResp,
		AccessToken: token,
	})
}

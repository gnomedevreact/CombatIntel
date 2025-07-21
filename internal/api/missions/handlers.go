package missions

import (
	"encoding/json"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"net/http"
)

type Handler struct {
	missionsService missionsService
	db              *database.Queries
	validator       *validator.Validate
}

func NewHandler(db *database.Queries, validator *validator.Validate) *Handler {
	return &Handler{
		missionsService: missionsService{
			db,
		},
		validator: validator,
		db:        db,
	}
}

func (h Handler) CreateMission() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var reqData Mission
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&reqData)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		var parsedMission Mission
		mission, err := h.missionsService.createMission(r.Context(), reqData)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		err = copier.Copy(&parsedMission, &mission)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, parsedMission)
	})
}

func (h Handler) GetAllMissions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var parsedMissions []Mission
		missions, err := h.db.GetAllMissions(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		if missions == nil {
			missions = []database.Mission{}
		}
		copier.Copy(&parsedMissions, &missions)
		utils.RespondWithJSON(w, http.StatusOK, parsedMissions)
	})
}

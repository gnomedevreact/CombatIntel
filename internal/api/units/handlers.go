package units

import (
	"encoding/json"
	"errors"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"net/http"
)

type Handler struct {
	unitsService unitsService
	validator    *validator.Validate
	db           *database.Queries
}

func NewHandler(db *database.Queries, validator *validator.Validate) *Handler {
	return &Handler{
		unitsService: unitsService{db},
		validator:    validator,
		db:           db,
	}
}

func (h *Handler) CreateUnit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var reqData createUnitRequest
	if err := decoder.Decode(&reqData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errors.New("an error occurred while processing your request"))
		return
	}

	if err := h.validator.Struct(reqData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	var parsedUnit Unit
	unit, err := h.unitsService.createUnit(r.Context(), reqData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = copier.Copy(&parsedUnit, unit)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, parsedUnit)
}

func (h *Handler) GetAllUnits(w http.ResponseWriter, r *http.Request) {
	var parsedUnits []Unit
	units, err := h.db.GetAllUnits(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = copier.Copy(&parsedUnits, &units)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, parsedUnits)
}

package missions

import (
	"context"
	"encoding/json"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCreateMission(t *testing.T) {
	ctx := context.Background()
	err := utils.LoadEnvFromProjectRoot()
	if err != nil {
		t.Fatal("Error loading .env file", err)
	}
	dbUrl := os.Getenv("TEST_DATABASE_URL")

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		t.Fatal("Error connecting to the database")
	}
	t.Log("connected to the database")
	queries := database.New(conn)
	validator := validator.New()

	reqBody := `{
  "name": "Alpha1",
  "objective": "Capture bridge",
  "start_time": "2025-07-21T08:00:00Z",
  "end_time": "2025-07-21T12:30:00Z",
  "outcome": "success",
  "losses": 12,
  "enemy_losses": 45,
  "enemy_forces_size": 100,
  "own_forces_size": 80,
  "notes": "Secured",
  "classification_level": "X",
  "unit_id": "19344f64-f850-4826-bb6c-5528882a9409"
}`
	req := httptest.NewRequest("POST", "/missions", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := NewHandler(queries, validator)
	handler.CreateMission().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var mission database.Mission
	err = json.NewDecoder(rr.Body).Decode(&mission)
	require.NoError(t, err)
	assert.Equal(t, "Alpha1", mission.Name)
	queries.DeleteMission(req.Context(), mission.ID)
}

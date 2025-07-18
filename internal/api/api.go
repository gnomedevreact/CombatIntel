package api

import (
	"github.com/gnomedevreact/CombatIntel/internal/database"
)

type ApiConfig struct {
	Db           *database.Queries
	Secret       string
	PublicSecret string
}

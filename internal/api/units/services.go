package units

import (
	"context"
	"github.com/gnomedevreact/CombatIntel/internal/database"
)

type unitsService struct {
	db *database.Queries
}

func (s *unitsService) createUnit(ctx context.Context, reqData createUnitRequest) (*database.Unit, error) {
	unit, err := s.db.CreateUnit(ctx, database.CreateUnitParams{
		Name:        reqData.name,
		CommanderID: reqData.commanderId,
	})
	if err != nil {
		return nil, err
	}
	return &unit, nil
}

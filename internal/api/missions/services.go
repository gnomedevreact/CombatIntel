package missions

import (
	"context"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/jinzhu/copier"
)

type missionsService struct {
	db *database.Queries
}

func (s *missionsService) createMission(ctx context.Context, reqData Mission) (*database.Mission, error) {
	var params database.CreateMissionParams
	err := copier.Copy(&params, &reqData)
	if err != nil {
		return nil, err
	}

	mission, err := s.db.CreateMission(ctx, params)
	if err != nil {
		return nil, err
	}

	return &mission, nil
}

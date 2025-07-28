package missions

import (
	"context"
	"errors"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
	"log"
	"strconv"
	"time"
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

func _parseRFC3339ToTz(s string) (pgtype.Timestamp, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return pgtype.Timestamp{}, errors.New("invalid time: " + s)
	}
	return pgtype.Timestamp{Time: t, Valid: true}, nil
}

func (s *missionsService) uploadMissions(ctx context.Context, records [][]string) (*[]Mission, error) {
	var missions []Mission
	for i, record := range records {
		if i == 0 {
			continue
		}

		createdAt, err := _parseRFC3339ToTz(record[1])
		if err != nil {
			return nil, err
		}
		updatedAt, err := _parseRFC3339ToTz(record[2])
		if err != nil {
			return nil, err
		}
		startTime, err := _parseRFC3339ToTz(record[4])
		if err != nil {
			return nil, err
		}
		endTime, err := _parseRFC3339ToTz(record[5])
		if err != nil {
			return nil, err
		}

		losses, err := strconv.Atoi(record[7])
		if err != nil {
			return nil, err
		}
		enemyLosses, err := strconv.Atoi(record[8])
		if err != nil {
			return nil, err
		}
		enemyForces, err := strconv.Atoi(record[9])
		if err != nil {
			return nil, err
		}
		ownForces, err := strconv.Atoi(record[10])
		if err != nil {
			return nil, err
		}

		var unitId pgtype.UUID
		if err := unitId.Scan(record[13]); err != nil {
			log.Println(err)
			return nil, err
		}

		mission := database.Mission{
			Name:                record[0],
			CreatedAt:           createdAt,
			UpdatedAt:           updatedAt,
			Objective:           record[3],
			StartTime:           startTime,
			EndTime:             endTime,
			Outcome:             record[6],
			Losses:              int32(losses),
			EnemyLosses:         int32(enemyLosses),
			EnemyForcesSize:     int32(enemyForces),
			OwnForcesSize:       int32(ownForces),
			Notes:               pgtype.Text{String: record[11], Valid: true},
			ClassificationLevel: record[12],
			UnitID:              unitId,
		}

		var params database.CreateMissionParams
		err = copier.Copy(&params, &mission)
		if err != nil {
			return nil, err
		}

		var parsedMission Mission
		mission, err = s.db.CreateMission(ctx, params)
		if err != nil {
			return nil, err
		}

		err = copier.Copy(&parsedMission, &mission)
		if err != nil {
			log.Println(err)
		}

		missions = append(missions, parsedMission)
	}

	return &missions, nil
}

func (s *missionsService) predictMissionResult(ctx context.Context, reqBody MissionPredict) (string, error) {
	missions, err := s.db.GetAllMissions(ctx)

	enemyForcesAttr := base.NewFloatAttribute("enemy_forces_size")
	ownForcesAttr := base.NewFloatAttribute("own_forces_size")
	outcomeAttr := base.NewCategoricalAttribute()

	attrs := []base.Attribute{enemyForcesAttr, ownForcesAttr, outcomeAttr}
	inst := base.NewDenseInstances()

	for _, attr := range attrs {
		inst.AddAttribute(attr)
	}
	inst.AddClassAttribute(outcomeAttr)
	enemyForcesSpec, err := inst.GetAttribute(enemyForcesAttr)
	ownForcesSpec, err := inst.GetAttribute(ownForcesAttr)
	outcomeSpec, err := inst.GetAttribute(outcomeAttr)

	inst.Extend(len(missions))
	for i, row := range missions {
		inst.Set(enemyForcesSpec, i, base.PackFloatToBytes(float64(row.EnemyForcesSize)))
		inst.Set(ownForcesSpec, i, base.PackFloatToBytes(float64(row.OwnForcesSize)))
		inst.Set(outcomeSpec, i, outcomeAttr.GetSysValFromString(row.Outcome))
	}

	tree := trees.NewID3DecisionTree(0.6)
	tree.Fit(inst)

	predictInst := base.NewDenseInstances()
	for _, attr := range attrs {
		predictInst.AddAttribute(attr)
	}
	predictInst.AddClassAttribute(outcomeAttr)
	predictInst.Extend(1)

	predictInst.Set(enemyForcesSpec, 0, base.PackFloatToBytes(float64(reqBody.EnemyForcesSize)))
	predictInst.Set(ownForcesSpec, 0, base.PackFloatToBytes(float64(reqBody.OwnForcesSize)))

	predictions, err := tree.Predict(predictInst)
	if err != nil {
		return "", err
	}

	return predictions.RowString(0), nil
}

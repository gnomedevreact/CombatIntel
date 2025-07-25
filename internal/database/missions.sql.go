// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: missions.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMission = `-- name: CreateMission :one
insert into missions
    (id,
     name,
     objective,
     start_time,
     end_time,
     outcome,
     losses,
     enemy_losses,
     enemy_forces_size,
     own_forces_size,
     notes,
     classification_level,
     unit_id)
values (
        gen_random_uuid(),
        $1,
        $2,
        $3,
        $4,$5,
        $6,$7,
        $8,$9,
        $10, $11,
        $12
) returning id, created_at, updated_at, name, objective, start_time, end_time, outcome, losses, enemy_losses, enemy_forces_size, own_forces_size, notes, classification_level, unit_id
`

type CreateMissionParams struct {
	Name                string
	Objective           string
	StartTime           pgtype.Timestamp
	EndTime             pgtype.Timestamp
	Outcome             string
	Losses              int32
	EnemyLosses         int32
	EnemyForcesSize     int32
	OwnForcesSize       int32
	Notes               pgtype.Text
	ClassificationLevel string
	UnitID              pgtype.UUID
}

func (q *Queries) CreateMission(ctx context.Context, arg CreateMissionParams) (Mission, error) {
	row := q.db.QueryRow(ctx, createMission,
		arg.Name,
		arg.Objective,
		arg.StartTime,
		arg.EndTime,
		arg.Outcome,
		arg.Losses,
		arg.EnemyLosses,
		arg.EnemyForcesSize,
		arg.OwnForcesSize,
		arg.Notes,
		arg.ClassificationLevel,
		arg.UnitID,
	)
	var i Mission
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Objective,
		&i.StartTime,
		&i.EndTime,
		&i.Outcome,
		&i.Losses,
		&i.EnemyLosses,
		&i.EnemyForcesSize,
		&i.OwnForcesSize,
		&i.Notes,
		&i.ClassificationLevel,
		&i.UnitID,
	)
	return i, err
}

const deleteMission = `-- name: DeleteMission :exec
delete from missions
where id = $1
`

func (q *Queries) DeleteMission(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteMission, id)
	return err
}

const getAllMissions = `-- name: GetAllMissions :many
select id, created_at, updated_at, name, objective, start_time, end_time, outcome, losses, enemy_losses, enemy_forces_size, own_forces_size, notes, classification_level, unit_id from missions
`

func (q *Queries) GetAllMissions(ctx context.Context) ([]Mission, error) {
	rows, err := q.db.Query(ctx, getAllMissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Mission
	for rows.Next() {
		var i Mission
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Objective,
			&i.StartTime,
			&i.EndTime,
			&i.Outcome,
			&i.Losses,
			&i.EnemyLosses,
			&i.EnemyForcesSize,
			&i.OwnForcesSize,
			&i.Notes,
			&i.ClassificationLevel,
			&i.UnitID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMissionById = `-- name: GetMissionById :one
select id, created_at, updated_at, name, objective, start_time, end_time, outcome, losses, enemy_losses, enemy_forces_size, own_forces_size, notes, classification_level, unit_id from missions
where id = $1
`

func (q *Queries) GetMissionById(ctx context.Context, id pgtype.UUID) (Mission, error) {
	row := q.db.QueryRow(ctx, getMissionById, id)
	var i Mission
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Objective,
		&i.StartTime,
		&i.EndTime,
		&i.Outcome,
		&i.Losses,
		&i.EnemyLosses,
		&i.EnemyForcesSize,
		&i.OwnForcesSize,
		&i.Notes,
		&i.ClassificationLevel,
		&i.UnitID,
	)
	return i, err
}

const getUnitMissions = `-- name: GetUnitMissions :many
select id, created_at, updated_at, name, objective, start_time, end_time, outcome, losses, enemy_losses, enemy_forces_size, own_forces_size, notes, classification_level, unit_id from missions
where unit_id = $1
`

func (q *Queries) GetUnitMissions(ctx context.Context, unitID pgtype.UUID) ([]Mission, error) {
	rows, err := q.db.Query(ctx, getUnitMissions, unitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Mission
	for rows.Next() {
		var i Mission
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Objective,
			&i.StartTime,
			&i.EndTime,
			&i.Outcome,
			&i.Losses,
			&i.EnemyLosses,
			&i.EnemyForcesSize,
			&i.OwnForcesSize,
			&i.Notes,
			&i.ClassificationLevel,
			&i.UnitID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Mission struct {
	ID                  pgtype.UUID
	CreatedAt           pgtype.Timestamp
	UpdatedAt           pgtype.Timestamp
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

type Unit struct {
	ID          pgtype.UUID
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	Name        string
	CommanderID pgtype.UUID
}

type User struct {
	ID             pgtype.UUID
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
	Username       string
	Password       string
	Role           string
	ClearanceLevel string
}

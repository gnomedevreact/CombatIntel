package missions

import "github.com/jackc/pgx/v5/pgtype"

type Mission struct {
	ID                  pgtype.UUID        `json:"id"`
	Name                string             `json:"name" validate:"required,alphanum"`
	CreatedAt           pgtype.Timestamptz `json:"created_at"`
	UpdatedAt           pgtype.Timestamptz `json:"updated_at"`
	Objective           string             `json:"objective" validate:"required"`
	StartTime           pgtype.Timestamptz `json:"start_time" validate:"required"`
	EndTime             pgtype.Timestamptz `json:"end_time" validate:"required"`
	Outcome             string             `json:"outcome" validate:"required"`
	Losses              int                `json:"losses" validate:"required,numeric"`
	EnemyLosses         int                `json:"enemy_losses" validate:"required,numeric"`
	EnemyForcesSize     int                `json:"enemy_forces_size" validate:"required,numeric"`
	OwnForcesSize       int                `json:"own_forces_size" validate:"required,numeric"`
	Notes               string             `json:"notes"`
	ClassificationLevel string             `json:"classification_level" validate:"required"`
	UnitId              pgtype.UUID        `json:"unit_id" validate:"required"`
}

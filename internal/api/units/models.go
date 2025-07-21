package units

import "github.com/jackc/pgx/v5/pgtype"

type Unit struct {
	ID          pgtype.UUID        `json:"id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	Name        string             `json:"name"`
	CommanderId pgtype.UUID        `json:"commander_id"`
}

type createUnitRequest struct {
	Name        string      `json:"name" validate:"required"`
	CommanderId pgtype.UUID `json:"commander_id" validate:"required"`
}

package units

import "github.com/jackc/pgx/v5/pgtype"

type createUnitRequest struct {
	Name        string      `json:"name" validate:"required"`
	CommanderId pgtype.UUID `json:"commander_id" validate:"required"`
}

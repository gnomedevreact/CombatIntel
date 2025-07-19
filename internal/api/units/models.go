package units

import "github.com/jackc/pgx/v5/pgtype"

type createUnitRequest struct {
	name        string      `json:"name" validate:"required"`
	commanderId pgtype.UUID `json:"commanderId" validate:"required"`
}

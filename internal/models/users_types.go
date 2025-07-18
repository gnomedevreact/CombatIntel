package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID             pgtype.UUID      `json:"id"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
	Username       string           `json:"username"`
	Role           string           `json:"role"`
	ClearanceLevel string           `json:"clearance_level"`
}

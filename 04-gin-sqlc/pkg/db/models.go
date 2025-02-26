// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID      `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Photo     string           `json:"photo"`
	Verified  bool             `json:"verified"`
	Password  string           `json:"password"`
	Role      string           `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

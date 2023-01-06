package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Campaign struct {
	id          pgtype.Int8
	createdAt   pgtype.Timestamptz
	name        pgtype.Text
	description pgtype.Text
}

type Campaigns struct {
	Campaigns []Campaign `json:"campaign"`
}

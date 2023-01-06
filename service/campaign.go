package service

import "github.com/jackc/pgx/v5/pgtype"

type Campaign struct {
	id          pgtype.Int4
	createdAt   pgtype.Timestamptz
	name        pgtype.Text
	description pgtype.Text
}

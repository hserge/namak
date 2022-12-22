package app

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

var db *pgx.Conn

func ConnectDb(ctx context.Context) {
	var err error
	// db, err = pgx.Connect(a.Ctx, dsn)
	if err != nil {
		log.Fatal().Err(err).Str("service", "App").Msgf("Unable to connect to database")
	}
}

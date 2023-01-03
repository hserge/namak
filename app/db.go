package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rs/zerolog/log"
)

func ConnectDb(ctx context.Context, dsn string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal().Err(err).Str("service", "Server").Msgf("Unable to connect to database")
	}

	return pool
}

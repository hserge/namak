package app

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Container struct {
	Ctx context.Context
	Db  *pgxpool.Pool
	Log zerolog.Logger
}

type GetDefaultParams struct {
	dsn string
}

func New(ctx context.Context, log zerolog.Logger, db *pgxpool.Pool) *Container {
	return &Container{
		Ctx: ctx,
		Db:  db,
		Log: log,
	}
}

func GetDefaults(params GetDefaultParams) (ctx context.Context, log zerolog.Logger, pool *pgxpool.Pool, err error) {
	if params.dsn == "" {
		params.dsn = os.Getenv("APP_DSN")
	}
	// context
	ctx = context.Background()

	// log
	log = zerolog.New(os.Stdout).With().Timestamp().Logger()

	// DB pool
	pool, err = pgxpool.New(ctx, params.dsn)
	if err != nil {
		log.Panic().Err(err).Str("service", "Server").Msgf("Unable to connect to database")
	}

	return ctx, log, pool, err
}

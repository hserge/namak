package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	Ctx context.Context
	Srv *echo.Echo
	Db  *pgx.Conn
	Log zerolog.Logger
}

func (a *App) Initialize(ctx context.Context, dsn string) {
	a.Ctx = ctx
	a.Srv = echo.New()

	a.SetMiddlewares()
	a.ConnectDb(dsn)
	a.InitializeRoutes()
}

func (a *App) SetMiddlewares() {
	// remove trailing slashes
	a.Srv.Pre(middleware.RemoveTrailingSlash())
	// set timeout for the request
	a.Srv.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
	// tag request with an id
	a.Srv.Use(middleware.RequestID())
	// gracefully recover
	a.Srv.Use(middleware.Recover())
	// dump body in log for debugging purposes
	// TODO: make sure this runs only in dev
	a.Srv.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		a.Log.Info().Interface("Request", reqBody).Interface("Response", resBody).Msg("BodyDumpLog")
	}))
	// det logger to zerolog
	a.Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	a.Srv.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			a.Log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
}

func (a *App) ConnectDb(dsn string) {
	var err error
	a.Db, err = pgx.Connect(a.Ctx, dsn)
	if err != nil {
		a.Log.Fatal().Err(err).Str("service", "App").Msgf("Unable to connect to database")
	}
}

func (a *App) CloseDb() {
	err := a.Db.Close(a.Ctx)
	if err != nil {
		a.Log.
			Fatal().
			Err(err).
			Str("service", "App").
			Msgf("Unable to close database connection")
	}
}

func (a *App) Start(addr string) {
	err := a.Srv.Start(addr)
	log.Fatal().Err(err).Str("service", "App").Msgf("Unable to run application")
}

func (a *App) InitializeRoutes() {
	a.Srv.Static("/", "static")

	a.Srv.GET("/", func(c echo.Context) error {
		var version, hostname string
		err := a.Db.QueryRow(a.Ctx, "SELECT VERSION() as version, pg_read_file('/etc/hostname') as hostname").Scan(&version, &hostname)
		if err != nil {
			// a.Log.Err(err)
			a.Log.Error().Msgf("Something is wrong: %v", err)
		}
		return c.String(http.StatusOK, "Hello World "+time.Now().Format(time.RFC3339)+"\nVersion:"+version+"\nHostname:"+hostname)
	})

	// a.Srv.POST("/emails", createEmailHandler)
	// a.Srv.GET("/emails/:id", readEmailHandler)
	// a.Srv.PUT("/emails/:id", updateEmailHandler)
	// a.Srv.DELETE("/emails/:id", deleteEmailHandler)
}

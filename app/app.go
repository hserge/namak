package app

import (
	"context"
	"github.com/edersohe/zflogger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hserge/namak/handler"
	"os"

	"github.com/elastic/go-sysinfo"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func Initialize(ctx context.Context) {
	// logger
	logger := GetLogger()
	// config
	InitConfig(logger)
	// db
	db := GetDb(ctx, logger, os.Getenv("PGSQL_DSN"))
	defer CloseDb(ctx, db, logger)

	app := fiber.New()

	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(zflogger.Middleware(logger, nil))

	ServerInfo(logger)
	InitializeRoutes(app)

	// Listen on port 3000
	err := app.Listen(os.Getenv("APP_PORT"))
	if err != nil {
		logger.Fatal().Err(err).Str("service", "Server").Msgf("Unable to run server on port: %s", os.Getenv("APP_PORT"))
	}
}

func ServerInfo(logger zerolog.Logger) {
	host, _ := sysinfo.Host()
	memory, _ := host.Memory()
	cpuTime, _ := host.CPUTime()

	logger.Info().
		Interface("Go", sysinfo.Go()).
		Interface("HostInfo", host.Info()).
		Interface("HostMemory", memory).
		Interface("HostCPUTime", cpuTime).
		Msg("System Information")
}

func InitConfig(logger zerolog.Logger) {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal().Err(err).Str("service", "Server").Msgf("Unable to load .env file")
	}
}

func GetLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func GetDb(ctx context.Context, logger zerolog.Logger, dsn string) *pgx.Conn {
	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		logger.Fatal().Err(err).Str("service", "Server").Msgf("Unable to connect to database")
	}

	// get the db version
	var version string
	err = db.QueryRow(ctx, "SELECT VERSION()").Scan(&version)
	if err != nil {
		logger.Fatal().Err(err).Str("service", "Server").Msgf("Unable to access database version")
	}

	logger.Info().Msgf("Connected to the database: %s", version)
	return db
}

func CloseDb(ctx context.Context, db *pgx.Conn, logger zerolog.Logger) {
	err := db.Close(ctx)
	if err != nil {
		logger.
			Fatal().
			Err(err).
			Str("service", "Server").
			Msgf("Unable to close database connection")
	}
}

func InitializeRoutes(app *fiber.App) {
	app.Static("/", "static")

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/emails", handler.EmailList)
	//v1.Post("/users", handlers.UserCreate)

	// Setup static files

	// s.Srv.POST("/emails", createEmailHandler)
	// s.Srv.GET("/emails/:id", readEmailHandler)
	// s.Srv.PUT("/emails/:id", updateEmailHandler)
	// s.Srv.DELETE("/emails/:id", deleteEmailHandler)

	v1.Get("/error", func(c *fiber.Ctx) error {
		a := 0
		return c.JSON(1 / a)
	})

	// Handle not founds
	app.Use(handler.NotFound)
}

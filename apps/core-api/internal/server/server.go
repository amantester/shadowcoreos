package server

import (
	"context"

	"github.com/amantester/shadowcoreos/apps/core-api/internal/config"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	cfg := config.Load()

	db, dbErr := database.Connect(cfg.DatabaseURL)

	app := fiber.New(fiber.Config{
		AppName: "ShadowCoreOS Core API",
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		dbStatus := "ok"

		if dbErr != nil || db == nil {
			dbStatus = "error"
		} else {
			if err := db.Ping(context.Background()); err != nil {
				dbStatus = "error"
			}
		}

		return c.JSON(fiber.Map{
			"status":   "ok",
			"service":  "shadowcoreos-core-api",
			"database": dbStatus,
			"env":      cfg.AppEnv,
		})
	})

	return app
}

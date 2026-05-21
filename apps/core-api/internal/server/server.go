package server

import (
	"context"

	"github.com/amantester/shadowcoreos/apps/core-api/internal/config"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/database"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	cfg := config.Load()

	db, dbErr := database.Connect(cfg.DatabaseURL)

	app := fiber.New(fiber.Config{
		AppName: "ShadowCoreOS Core API",
	})

	app.Use(middleware.RequestID())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", func(c *fiber.Ctx) error {
		dbStatus := "ok"

		if dbErr != nil || db == nil {
			dbStatus = "error"
		} else {
			if err := db.Ping(context.Background()); err != nil {
				dbStatus = "error"
			}
		}

		return c.JSON(fiber.Map{
			"status":     "ok",
			"service":    "shadowcoreos-core-api",
			"database":   dbStatus,
			"env":        cfg.AppEnv,
			"request_id": c.Locals("request_id"),
		})
	})

	return app
}

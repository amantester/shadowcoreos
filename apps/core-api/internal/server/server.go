package server

import (
	"context"

	"github.com/amantester/shadowcoreos/apps/core-api/internal/auth"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/config"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/database"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/handlers"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/middleware"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/repositories"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	cfg := config.Load()

	db, dbErr := database.Connect(cfg.DatabaseURL)

	app := fiber.New(fiber.Config{
		AppName: "ShadowCoreOS Core API",
	})

	app.Use(middleware.RequestID())

	userRepo := repositories.NewUserRepository(db)
authService := services.NewAuthService(userRepo)
authHandler := handlers.NewAuthHandler(authService)

projectRepo := repositories.NewProjectRepository(db)
projectService := services.NewProjectService(projectRepo)
projectHandler := handlers.NewProjectHandler(projectService)

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

	authRoutes := v1.Group("/auth")

	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)

	projectRoutes := v1.Group("/projects")
projectRoutes.Use(auth.Protected())

projectRoutes.Post("/", projectHandler.Create)
projectRoutes.Get("/", projectHandler.List)

	protected := v1.Group("/protected")

	protected.Use(auth.Protected())

	protected.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "authenticated access granted",
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})

	admin := v1.Group("/admin")

	admin.Use(auth.Protected())
	admin.Use(auth.RequireRole("admin"))

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "admin access granted",
		})
	})

	return app
}
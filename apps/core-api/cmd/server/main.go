package main

import (
	"log"

	"github.com/amantester/shadowcoreos/apps/core-api/internal/config"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/server"
)

func main() {
	cfg := config.Load()
	app := server.New()

	log.Printf("ShadowCoreOS Core API running on :%s", cfg.Port)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

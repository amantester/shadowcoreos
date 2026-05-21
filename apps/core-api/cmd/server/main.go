package main

import (
	"log"

	"github.com/amantester/shadowcoreos/apps/core-api/internal/server"
)

func main() {
	app := server.New()

	log.Println("ShadowCoreOS Core API running on :8080")

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

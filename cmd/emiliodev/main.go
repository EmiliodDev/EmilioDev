package main

import (
	"log"

	"github.com/EmiliodDev/EmilioDev/config"
	"github.com/EmiliodDev/EmilioDev/internal/db"
	"github.com/EmiliodDev/EmilioDev/internal/router"
)

func main() {
	config.LoadConfig()

	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("connection db failed, %v", err)
	}

	r := router.SetupRouter()

	if err := r.Run(":" + config.AppConfig.Server.Port); err != nil {
		log.Fatalf("error starting the server, %v", err.Error())
	}
}

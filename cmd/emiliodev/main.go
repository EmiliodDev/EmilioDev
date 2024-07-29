package main

import (
	"log"

	"github.com/EmiliodDev/EmilioDev/config"
	"github.com/EmiliodDev/EmilioDev/internal/db"
)

func main() {
    config.LoadConfig()

    if db, err := db.ConnectToDB(); err != nil {
        log.Fatalf("db connection failed, %v", err)
    } 


}

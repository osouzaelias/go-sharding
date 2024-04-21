package main

import (
	"go-sharding/config"
	"go-sharding/internal/adapters/db"
	"go-sharding/internal/adapters/rest"
	"go-sharding/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetRegion())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, config.GetTables())
	restAdapter := rest.NewAdapter(application, config.GetPort())
	restAdapter.Run()
}

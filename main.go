package main

import (
	"log"
	"os"

	"github.com/h00s/bitbox/api"
	"github.com/h00s/bitbox/api/config"
	"github.com/h00s/bitbox/api/db"
)

func main() {
	config := config.NewConfig()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	database := db.NewDatabase(&config.Database)
	if err := database.Connect(); err != nil {
		logger.Fatal(err)
	}
	if err := database.Migrate(); err != nil {
		logger.Fatal(err)
	}
	defer database.Close()

	api := api.NewAPI(config, database, logger)
	api.Start()
	api.WaitForShutdown()
}

package main

import (
	"back/internal/delivery"
	"back/pkg/config"
	"back/pkg/database"
	"back/pkg/log"
)

func main() {
	log, loggerInfoFile, loggerErrorFile := log.InitLogger()

	defer loggerInfoFile.Close()
	defer loggerErrorFile.Close()

	config.InitConfig()
	log.Info("Config initialized")

	db := database.GetDB()
	log.Info("Database initialized")

	delivery.Start(db, log)

}

package main

import (
	"go-backend/internal/config"
	"go-backend/internal/database"
	"go-backend/internal/util"
	"go-backend/pkg/auth"
	// other imports
)

func main() {
	// Initialize logger
	util.InitLogger()
	// Initialize configuration
	config.InitConfig()
	// now we need to connect to postgres
	errPostgres := util.Retry(func() error {
		return database.NewPostgresDBPool(&config.Config.Database)
	})
	if errPostgres != nil {
		util.Logger.Sugar().Fatalf("Error connecting to PostgresSQL: %s", errPostgres)
	}

	// Initialize features
	if config.Config.Features.EnableAuth {
		auth.InitAuth()
	}

	util.Logger.Sugar().Info("Application running!")
	select {
	// do nothing
	}
}

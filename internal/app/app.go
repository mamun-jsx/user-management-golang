package app

import (
	"log"

	"github.com/mamun-jsx/user-management-golang/config"
	"github.com/mamun-jsx/user-management-golang/internal/database"
	"gorm.io/gorm"
)

// define the apps type | blue print
type App struct {
	Config *config.Config
	DB     *gorm.DB
}

// make an function to call anywhere or mount

func BootstrapApp() *App {
	//* load all Config
	cfg := config.LoadEnv()

	// pass the config file to database to connect
	db, err := database.DatabaseConnection(cfg)
	if err != nil {
		log.Fatal("unable to connect database from app", err.Error())
	}

	//? return and store the value into App
	return &App{
		Config: cfg,
		DB:     db,
	}
}

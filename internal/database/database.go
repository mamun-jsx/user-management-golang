package database

import (
	"fmt"
	"log"

	"github.com/mamun-jsx/user-management-golang/config"
	"github.com/mamun-jsx/user-management-golang/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection(cfg *config.Config) (*gorm.DB, error) {
	// TODO 1 : Create Database Link
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)

	// TODO 2 : open database by gorm
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	// TODO 3 : handle error if database unable to open
	if err != nil {
		log.Fatal("Unable to connect database")
	}
	// TODO 4 : mount auto-migration into database with model
	if err := db.AutoMigrate(&models.ProductModel{}, &models.UserModel{}); err != nil {
		return nil, fmt.Errorf("unable to migrate database: %w", err)
	}
	// TODO 5 : return database and nil
	log.Printf("🚀🚀🚀Database Connected Successfully🚀🚀🚀")
	return db, nil
}

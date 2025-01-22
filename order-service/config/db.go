package config

import (
	"log"

	"ecommerce-microservices/common/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(appConfig *config.AppConfig) (*gorm.DB, error) {
	dsn := appConfig.Database.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database successfully.")
	return db, nil
}

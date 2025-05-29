package main

import (
	"auth-service/helper"
	my_logger "auth-service/logger"
	"auth-service/models"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// RUN MIGRATE TABLES
/*
go run "/home/spionera/Documents/me/code/go_clean_arch+msvcs/job-queue-auth-service/migration/main.go"
*/
func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error loading .env file: %v", err)
	}

	PSQL_CONNECTION := helper.GetENV("PSQL_CONNECTION", "")
	db, err := gorm.Open(postgres.Open(PSQL_CONNECTION), &gorm.Config{
		Logger: my_logger.GormLogger(),
	})
	if err != nil {
		logrus.Errorf("Error connecting to database: %v", err)
	}
	if err := db.Migrator().AutoMigrate(
		&models.Account{},
		&models.Session{},
		&models.UserOTP{},
		&models.UserConsent{},
		&models.Consent{},
	); err != nil {
		logrus.Errorf("Error migrating database table: %v", err)
	}

}

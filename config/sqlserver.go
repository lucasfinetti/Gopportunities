package config

import (
	"fmt"

	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server   = "dpg-cvntfsk9c44c73fg9n10-a"
	port     = 5432
	user     = "gopportunities_db_user"
	password = "8gRZYmsZiEWyHzPMXXsSQYYJy5DOaaIu"
	database = "gopportunities_db"
)

func InitializePostgres() (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
		server,
		user,
		password,
		database,
		port,
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}

	models := []interface{}{
		&schemas.Opening{},
		&schemas.Candidate{},
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		fmt.Printf("PostgreSQL automigration error: %v\n", err)
		return nil, err
	}

	return db, nil
}

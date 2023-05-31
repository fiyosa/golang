package cmd

import (
	"log"
	"tutorial/config/db"
	"tutorial/database/migration"
)

func InitMigrate(status bool) {
	err := db.DB.AutoMigrate(migration.Models()...)

	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
		return
	}

	log.Println("Migrate table successfully")
}

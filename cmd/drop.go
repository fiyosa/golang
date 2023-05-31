package cmd

import (
	"log"
	"tutorial/config/db"
	"tutorial/database/migration"
)

func InitDrop() {
	err := db.DB.Migrator().DropTable(migration.Models()...)

	if err != nil {
		log.Fatalf("Failed drop all table: %v", err)
		return
	}

	log.Print("Deleted all table successfully")
}

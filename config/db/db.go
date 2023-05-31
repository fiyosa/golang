package db

import (
	"fmt"
	"log"
	"tutorial/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(status bool) *gorm.DB {
	dsn := fmt.Sprintf(`host=%v user=%v password="%v" dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta`,
		env.ENV_DB_HOST,
		env.ENV_DB_USER,
		env.ENV_DB_PASS,
		env.ENV_DB_NAME,
		env.ENV_DB_PORT,
	)

	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		// Logger:                 logger.Default.LogMode(logger.Info), // show script query sql
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	if status {
		log.Println("Connected database successful")
	}

	DB = connect

	return connect
}

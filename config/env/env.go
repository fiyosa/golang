package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var ENV_PORT string
var ENV_DEBUG bool

var ENV_DB_HOST string
var ENV_DB_PORT string
var ENV_DB_NAME string
var ENV_DB_USER string
var ENV_DB_PASS string

func InitEnv() bool {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return false
	}

	if debug, err := strconv.ParseBool(os.Getenv("ENV_DEBUG")); err == nil {
		ENV_DEBUG = debug
	}

	if debug, err := strconv.ParseBool(os.Getenv("ENV_DEBUG")); err == nil {
		ENV_DEBUG = debug
	} else {
		ENV_DEBUG = true
	}

	getEnv()
	return true
}

func getEnv() {
	ENV_PORT = ":" + os.Getenv("ENV_PORT")

	ENV_DB_HOST = os.Getenv("ENV_DB_HOST")
	ENV_DB_PORT = os.Getenv("ENV_DB_PORT")
	ENV_DB_NAME = os.Getenv("ENV_DB_NAME")
	ENV_DB_USER = os.Getenv("ENV_DB_USER")
	ENV_DB_PASS = os.Getenv("ENV_DB_PASS")
}

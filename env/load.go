package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	POSTGRES_URL string
	BACKUP_CRON  string
}

var ENVIRONMENT = Environment{
	POSTGRES_URL: "",
	BACKUP_CRON:  "",
}

func loadStringFromEnv(key string) string {
	envVar := os.Getenv(key)

	if envVar == "" {
		panic("Missing environment variable: " + key)
	}

	return envVar
}

func loadIntoEnvironment() {
	ENVIRONMENT.POSTGRES_URL = loadStringFromEnv("POSTGRES_URL")
	ENVIRONMENT.BACKUP_CRON = loadStringFromEnv("BACKUP_CRON")
}

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	loadIntoEnvironment()
}

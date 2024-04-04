package env

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Environment struct {
	POSTGRES_URL string
	BACKUP_CRON  string
}

func (e *Environment) GetKeys() []string {
	var keys []string
	t := reflect.TypeOf(*e)

	for i := 0; i < t.NumField(); i++ {
		keys = append(keys, t.Field(i).Name)
	}

	return keys
}

func (e *Environment) SetValue(key string, value string) {
	rv := reflect.ValueOf(e)
	rv = rv.Elem()
	fv := rv.FieldByName(key)
	fv.SetString(value)
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
	for _, key := range ENVIRONMENT.GetKeys() {
		ENVIRONMENT.SetValue(key, loadStringFromEnv(key))
	}
}

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	loadIntoEnvironment()
}

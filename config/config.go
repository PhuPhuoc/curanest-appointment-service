package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort          string
	EnvDev           string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	Key              string
	PayOsClientId    string
	PayOsApiKey      string
	PayOsCheckSumKey string
	GoongAPIURL      string
	GoongAPIKEY      string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found => using system environment variables")
	}

	AppConfig = Config{
		AppPort:          ":" + os.Getenv("APP_PORT"),
		EnvDev:           os.Getenv("ENV_DEV"),
		DBHost:           os.Getenv("DB_HOST"),
		DBPort:           os.Getenv("DB_PORT"),
		DBUser:           os.Getenv("DB_USER"),
		DBPassword:       os.Getenv("DB_PASSWORD"),
		DBName:           os.Getenv("DB_NAME"),
		Key:              os.Getenv("KEY"),
		PayOsClientId:    os.Getenv("PAYOS_CLIENTID"),
		PayOsApiKey:      os.Getenv("PAYOS_APIKEY"),
		PayOsCheckSumKey: os.Getenv("PAYOS_CHECKSUMKEY"),
		GoongAPIURL:      os.Getenv("GOONG_API_URL"),
		GoongAPIKEY:      os.Getenv("GOONG_API_KEY"),
	}
}

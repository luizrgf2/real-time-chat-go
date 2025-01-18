package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DbHost     *string
	DbUser     *string
	DbPass     *string
	DbPort     *int
	DbDatabase *string
}

func LoadDatabaseConfig() DatabaseConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	DbDatabase := os.Getenv("DB_DATABASE")

	config := DatabaseConfig{
		DbHost:     &DbHost,
		DbUser:     &DbUser,
		DbPass:     &DbPass,
		DbPort:     &DbPort,
		DbDatabase: &DbDatabase,
	}
	return config
}

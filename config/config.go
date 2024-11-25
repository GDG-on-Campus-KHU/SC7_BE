package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	Port       int
}

func New() (*Config, error) {
	port := os.Getenv("PORT")
	dbPort := os.Getenv("DB_PORT")

	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     atoiOrDefault(dbPort, 3306),
		DBName:     os.Getenv("DB_NAME"),
		Port:       atoiOrDefault(port, 8080),
	}, nil
}

func atoiOrDefault(value string, defaultValue int) int {
	if v, err := strconv.Atoi(value); err == nil {
		return v
	}
	return defaultValue
}

package database

import (
	"fmt"
	"os"
)

// Строка подключения к базе данных
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func getConnString() string {
	user := getEnv("SHORTER_DB_USER", "user")
	password := getEnv("SHORTER_DB_PASSWORD", "password")
	hostname := getEnv("SHORTER_DB_HOSTNAME", "localhost")
	port := getEnv("SHORTER_DB_PORT", "5432")
	dbName := getEnv("SHORTER_DB_NAME", "postgres")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, hostname, port, dbName)
	return connString
}

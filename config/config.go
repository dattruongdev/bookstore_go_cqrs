package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	DriverName    string
	DataSource    string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassword    string
	DbName        string
	SSL_Mode      string
	MIGRATION_DIR string
}

func NewConfig() *Config {
	driverName := os.Getenv("DRIVER_NAME")
	dataSource := os.Getenv("DATA_SOURCE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")
	migration_dir := os.Getenv("MIGRATION_DIR")

	return &Config{
		DriverName:    driverName,
		DataSource:    dataSource,
		DbHost:        dbHost,
		DbPort:        dbPort,
		DbUser:        dbUser,
		DbPassword:    dbPassword,
		DbName:        dbName,
		SSL_Mode:      sslMode,
		MIGRATION_DIR: migration_dir,
	}
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName, c.SSL_Mode,
	)
}

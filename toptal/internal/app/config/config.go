package config

import "os"

type Config struct {
	HTTPAddr       string
	DSN            string
	MigrationsPath string
}

func Read() Config {
	var config Config
	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HTTPAddr = httpAddr
	}
	dsn, exists := os.LookupEnv("DSN")
	if exists {
		config.DSN = dsn
	}
	migrationsPath, exists := os.LookupEnv("MIGRATIONS_PATH")
	if exists {
		config.MigrationsPath = migrationsPath
	}
	return config
}

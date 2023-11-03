package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	DatabaseDSN string // DatabaseDSN PostgreSQL data source name.
	HostAddr    string // Host server's address.
	LogLevel    string // log level.
}

func Parse() Config {
	var config = Config{}
	checkFlags(&config)
	checkEnvironments(&config)

	return config
}

const (
	flagHostAddress = "a"
	flagDatabaseDSN = "d"
	flagLogLevel    = "l"
)

// checkFlags checks flags of app's launch.
func checkFlags(config *Config) {
	// main app.
	flag.StringVar(&config.HostAddr, flagHostAddress, "localhost:8080", "server endpoint")

	// postgres.
	flag.StringVar(&config.DatabaseDSN, flagDatabaseDSN, "postgres://postgres:postgres@localhost:5432/revtracker_test?sslmode=disable", "database DSN")

	// log.
	flag.StringVar(&config.LogLevel, flagLogLevel, "info", "log level")

	flag.Parse()
}

type envConfig struct {
	DatabaseDSN string `env:"DATABASE_URI"`
	HostAddr    string `env:"RUN_ADDRESS"`
	LogLevel    string `env:"LOG_LEVEL"`
}

// checkEnvironments checks environments suitable for server.
func checkEnvironments(config *Config) {
	var envs = envConfig{}
	err := env.Parse(&envs)
	if err != nil {
		log.Fatal(err)
	}

	// main app.
	_ = SetEnvToParamIfNeed(&config.HostAddr, envs.HostAddr)

	// postgres.
	_ = SetEnvToParamIfNeed(&config.DatabaseDSN, envs.DatabaseDSN)

	//log level.
	_ = SetEnvToParamIfNeed(&config.LogLevel, envs.LogLevel)
}

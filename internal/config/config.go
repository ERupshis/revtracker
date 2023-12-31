package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
	"github.com/erupshis/revtracker/internal/config/utils"
)

type Config struct {
	DatabaseDSN      string // DatabaseDSN PostgreSQL data source name
	DatabaseIdleConn int
	DatabaseOpenConn int
	HostAddr         string // Host server's address
	JWTKey           string // jwt web token generation key
	JWTExpiration    int    // jwt web token expiration time (hours)
	LogLevel         string // log level
}

func Parse() Config {
	var config = Config{}
	checkFlags(&config)
	checkEnvironments(&config)

	return config
}

const (
	flagHostAddress      = "a"
	flagDatabaseDSN      = "d"
	flagLogLevel         = "l"
	flagDatabaseIdleConn = "di"
	flagDatabaseOpenConn = "do"
	flagJWTKey           = "j"
	flagJWTExpiration    = "je"
)

// checkFlags checks flags of app's launch.
func checkFlags(config *Config) {
	// main app.
	flag.StringVar(&config.HostAddr, flagHostAddress, "localhost:8080", "server endpoint")

	// auth.
	flag.StringVar(&config.JWTKey, flagJWTKey, "TO REMOVE DEFAULT", "JWT web token key")
	flag.IntVar(&config.JWTExpiration, flagJWTExpiration, 24, "JWT expiration time")

	// postgres.
	flag.StringVar(&config.DatabaseDSN, flagDatabaseDSN, "postgres://postgres:postgres@localhost:5432/revtracker_test?sslmode=disable", "database DSN")
	flag.IntVar(&config.DatabaseIdleConn, flagDatabaseIdleConn, 3, "database max idle connections")
	flag.IntVar(&config.DatabaseOpenConn, flagDatabaseOpenConn, 3, "database max open connections")

	// log.
	flag.StringVar(&config.LogLevel, flagLogLevel, "info", "log level")

	flag.Parse()
}

type envConfig struct {
	DatabaseDSN      string `env:"DB_DSN"`
	DatabaseIdleConn string `env:"DB_MAX_IDLE_CONN"`
	DatabaseOpenConn string `env:"DB_MAX_OPEN_CONN"`
	HostAddr         string `env:"RUN_ADDRESS"`
	JWTKey           string `env:"JWT_KEY"`
	JWTExpiration    string `env:"JWT_EXPIRATION"`
	LogLevel         string `env:"LOG_LEVEL"`
}

// checkEnvironments checks environments suitable for server.
func checkEnvironments(config *Config) {
	var envs = envConfig{}
	err := env.Parse(&envs)
	if err != nil {
		log.Fatal(err)
	}

	// main app.
	_ = utils.SetEnvToParamIfNeed(&config.HostAddr, envs.HostAddr)

	// authentication.
	_ = utils.SetEnvToParamIfNeed(&config.JWTKey, envs.JWTKey)
	_ = utils.SetEnvToParamIfNeed(&config.JWTExpiration, envs.JWTExpiration)
	// postgres.
	_ = utils.SetEnvToParamIfNeed(&config.DatabaseDSN, envs.DatabaseDSN)
	_ = utils.SetEnvToParamIfNeed(&config.DatabaseIdleConn, envs.DatabaseIdleConn)
	_ = utils.SetEnvToParamIfNeed(&config.DatabaseOpenConn, envs.DatabaseOpenConn)

	// log level.
	_ = utils.SetEnvToParamIfNeed(&config.LogLevel, envs.LogLevel)
}

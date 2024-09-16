package config

import (
	"log"
)

const (
	LocalEnv = "local"
	ProdEnv  = "prod"

	env           = "ENV"
	serverAddress = "SERVER_ADDRESS"
	postgresConn  = "POSTGRES_CONN"
)

// Config is a structure that holds the application configuration.
type Config struct {
	Env           string
	ServerAddress string
	PostgresConn  string
}

// MustLoad reads the configuration from the environment variables.
func MustLoad(getenv func(string) (string, bool)) *Config {
	mustGet := func(key string) string {
		val, ok := getenv(key)
		if !ok {
			log.Fatalf("missing %s", key)
		}
		return val
	}

	servAddr := mustGet(serverAddress)
	pgsConn := mustGet(postgresConn)
	env, ok := getenv(env)
	if !ok {
		env = ProdEnv
	}

	return &Config{
		Env:           env,
		ServerAddress: servAddr,
		PostgresConn:  pgsConn,
	}
}

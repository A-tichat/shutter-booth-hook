package config

import (
	"os"
)

type Config struct {
	Env  string     `required:"true"`
	HTTP HTTPConfig `required:"true"`
}

const defaultEnv = "development"

func NewConfig() *Config {
	env := defaultEnv
	if goEnv := os.Getenv("GO_ENV"); goEnv != "" {
		env = goEnv
	}

	httpConfig := LoadHTTPConfig()

	return &Config{
		Env:  env,
		HTTP: httpConfig,
	}

}

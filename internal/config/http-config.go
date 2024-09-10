package config

import (
	"fmt"
	"os"
)

type HTTPConfig struct {
	Host string `required:"true"`
	Port string `required:"true"`
}

const defaultPort = "3000"

func LoadHTTPConfig() HTTPConfig {
	port := defaultPort
	if appPort := os.Getenv("PORT"); appPort != "" {
		port = appPort
	}

	return HTTPConfig{
		Host: os.Getenv("HOST"),
		Port: port,
	}
}

func (c HTTPConfig) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

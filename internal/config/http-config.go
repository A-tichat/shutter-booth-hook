package config

import (
	"fmt"
	"os"
)

type HTTPConfig struct {
	Host string `required:"true"`
	Port string `required:"true"`
}

const defaultPort = "8080"

func LoadHTTPConfig() HTTPConfig {
	port := defaultPort
	if goPort := os.Getenv("GO_PORT"); goPort != "" {
		port = goPort
	}

	return HTTPConfig{
		Host: os.Getenv("HOST"),
		Port: port,
	}
}

func (c HTTPConfig) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

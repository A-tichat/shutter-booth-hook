package config

import "os"

type OmiseConfig struct {
	ApiKey    string `required:"true"`
	SecretKey string `required:"true"`
}

func LoadOmiseConfig() OmiseConfig {
	return OmiseConfig{
		ApiKey:    os.Getenv("OMISE_API_KEY"),
		SecretKey: os.Getenv("OMISE_SECRET_KEY"),
	}
}

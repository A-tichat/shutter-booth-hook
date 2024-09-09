package main

import (
	"github.com/a-tichat/go-web/internal/config"
	"github.com/a-tichat/go-web/internal/server"
)

func main() {
	config := config.NewConfig()

	server.Server(config)
}

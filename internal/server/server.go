package server

import (
	"fmt"

	"git.akyoto.dev/go/web"
	"github.com/a-tichat/go-web/internal/commons/router"
	"github.com/a-tichat/go-web/internal/config"
)

func Server() {
	s := web.NewServer()

	router.ConfigureRoutes(s)
	config := config.NewConfig()

	fmt.Printf("Server running \"%s\" mode on port %s\n", config.Env, config.HTTP.Port)
	s.Run(config.HTTP.GetServerAddr())
}

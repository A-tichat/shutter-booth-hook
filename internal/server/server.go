package server

import (
	"fmt"

	"git.akyoto.dev/go/web"
	"github.com/a-tichat/go-web/internal/commons/router"
	"github.com/a-tichat/go-web/internal/config"
)

func Server(conf *config.Config) {
	s := web.NewServer()

	router.ConfigureRoutes(s)

	fmt.Printf("Server running \"%s\" mode on port %s\n", conf.Env, conf.HTTP.Port)
	s.Run(conf.HTTP.GetServerAddr())
}

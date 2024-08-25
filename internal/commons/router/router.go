package router

import (
	"fmt"
	"net/http"
	"time"

	"git.akyoto.dev/go/web"
	"github.com/a-tichat/go-web/internal/health"
	"github.com/a-tichat/go-web/internal/omise"
)

func ConfigureRoutes(s web.Server) {
	router := s.Router()

	healthController := health.New()
	router.Add(http.MethodGet, "/health", healthController.CheckHealth)

	omiseController := omise.New()
	router.Add(http.MethodPost, "/omise-callback", omiseController.Callback)

	s.Use(func(c web.Context) error {
		start := time.Now()

		defer func() {
			fmt.Printf("[%s] %s %s %d - %v\n",
				time.Now().Format("2006-01-02 15:04:05"),
				c.Request().Method(),
				c.Request().Path(),
				c.Response().Status(),
				time.Since(start))
		}()

		return c.Next()
	})

}

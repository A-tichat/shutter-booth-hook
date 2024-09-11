package server

import (
	"fmt"
	"net/http"
	"time"

	"git.akyoto.dev/go/web"
	"github.com/a-tichat/go-web/internal/config"
	"github.com/a-tichat/go-web/internal/health"
	"github.com/a-tichat/go-web/internal/omise"
)

const HEALTH_PATH = "/health"
const METRICS_PATH = "/metrics"

func Server(conf *config.Config) {
	s := web.NewServer()

	IGNORE_PATH := getIgnorePath()
	router := s.Router()

	healthCtr := health.New()
	router.Add(http.MethodGet, "/health", healthCtr.CheckHealth)

	omiseCtr := omise.New()
	router.Add(http.MethodPost, "/omise-callback", omiseCtr.Callback)

	s.Use(func(c web.Context) error {
		start := time.Now()

		defer func() {
			if _, ok := IGNORE_PATH[c.Request().Path()]; ok {
				return
			}

			fmt.Printf("[%s] %s %s %d - %v\n",
				time.Now().Format("2006-01-02 15:04:05"),
				c.Request().Method(),
				c.Request().Path(),
				c.Response().Status(),
				time.Since(start))
		}()

		return c.Next()
	})

	fmt.Printf("Server running \"%s\" mode on port %s\n", conf.Env, conf.HTTP.Port)
	s.Run(conf.HTTP.GetServerAddr())
}

// get path to ignore logging
func getIgnorePath() map[string]bool {
	fmt.Println("GET IGNORE PATH")
	return map[string]bool{
		HEALTH_PATH:  true,
		METRICS_PATH: true,
	}
}

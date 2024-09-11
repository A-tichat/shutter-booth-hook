package server

import (
	"context"
	"fmt"
	"time"

	"github.com/a-tichat/go-web/internal/config"
	healthCtr "github.com/a-tichat/go-web/internal/health"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const HEALTH_PATH = "/health"
const METRICS_PATH = "/metrics"

func Server(conf *config.Config) {
	IGNORE_PATH := getIgnorePath()
	hlog.SetLevel(hlog.LevelDebug)

	s := server.Default(
		server.WithHostPorts(conf.HTTP.GetServerAddr()),
	)

	s.Use(func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()

		defer func() {
			if _, ok := IGNORE_PATH[string(c.Path())]; ok {
				return
			}

			hlog.Debugf("%s %s %d - %v\n",
				c.Method(),
				c.Path(),
				c.Response.StatusCode(),
				time.Since(start))

			// fmt.Printf("[%s] %s %s %d - %v\n",
			// 	time.Now().Format("2006-01-02 15:04:05"),
			// 	c.Request().Method(),
			// 	c.Request().Path(),
			// 	c.Response().Status(),
			// 	time.Since(start))
		}()

		c.Next(ctx)
	})

	s.GET("/health", healthCtr.CheckHealth)

	s.Spin()
}

// get path to ignore logging
func getIgnorePath() map[string]bool {
	fmt.Println("GET IGNORE PATH")
	return map[string]bool{
		HEALTH_PATH:  true,
		METRICS_PATH: true,
	}
}

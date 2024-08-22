package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"git.akyoto.dev/go/web"
)

func main() {
	s := web.NewServer()
	router := s.Router()

	router.Add(http.MethodGet, "/health", func(ctx web.Context) error {
		return ctx.String("OK")
	})

	router.Add(http.MethodPost, "/omise-webhook", func(ctx web.Context) error {
		return ctx.String("Success")
	})

	s.Use(func(ctx web.Context) error {
		start := time.Now()

		defer func() {
			fmt.Println(ctx.Request().Path(), time.Since(start))
		}()

		return ctx.Next()
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddr := fmt.Sprintf(":%s", port)

	fmt.Println("Server running on http://localhost" + serverAddr)
	s.Run(serverAddr)
}

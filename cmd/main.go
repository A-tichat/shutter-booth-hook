package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler func(rw http.ResponseWriter, req *http.Request) error

type Server interface {
	Run(address string) error
	Use(handlers ...Handler)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			if r.URL.Path == "/health" {
				return
			}
			fmt.Printf("[%s] %s %s - %v\n",
				time.Now().Format("2006-01-02 15:04:05"),
				r.Method,
				r.URL.Path,
				time.Since(start))
		}()

		switch r.Method {
		case http.MethodPost:
			omiseReceiver(w, r)
			fmt.Fprintf(w, "Hello World!")

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})

	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type OmiseData struct {
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}

func omiseReceiver(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body OmiseData
	if err := decoder.Decode(&body); err != nil {
		panic(err)
	}

	fmt.Printf("Key: %s, Data: %v\n", body.Key, body.Data)

}

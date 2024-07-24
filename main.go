package main

import (
	"log"
	"net/http"
	"time"

	"github.com/willchat-ofc/api-willchat-golang/internal/setup"
	"github.com/willchat-ofc/api-willchat-golang/internal/setup/config"
)

func main() {
	port := ":7070"
	config.LoadEnvFile(".env")

	log.Println("server is running with port", port)

	sm := http.Server{
		Addr:         port,
		Handler:      setup.Server(),
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	sm.ListenAndServe()
}

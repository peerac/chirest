package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"peerac/go-chi-rest-example/config"
)

type application struct {
	config config.AppConfig
}

func main() {
	var cfg config.AppConfig
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (dev|stag|prod")
	flag.StringVar(&cfg.Version, "version", "1.0.0", "Running application version")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
	}

	// create the server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	logger.Printf("starting %s server on %s", cfg.Env, server.Addr)
	err := server.ListenAndServe()
	logger.Fatal(err)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"peerac/go-chi-rest-example/config"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	config.LoadEnv()
}

type application struct {
	config config.AppConfig
}

func main() {
	var cfg config.AppConfig

	// load env variables
	appPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	flag.IntVar(&cfg.Port, "port", appPort, "API server port")
	flag.StringVar(&cfg.Env, "env", os.Getenv("APP_ENV"), "Environment (dev|stag|prod")
	flag.StringVar(&cfg.Version, "version", os.Getenv("APP_VERSION"), "Running application version")
	flag.StringVar(&cfg.Database.DSN, "db-dsn", os.Getenv("DSN"), "database DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
	}

	// create the server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(db),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	logger.Printf("starting %s server on %s", cfg.Env, server.Addr)
	err = server.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.DSN), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	db.WithContext(ctx)

	return db, nil
}

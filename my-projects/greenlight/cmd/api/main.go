package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"greenlight/internal/data"
	postgres "greenlight/pkg/postgres"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		connectionString string
	}
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// If you’re following along, create a new GREENLIGHT_DB_DSN(connection string)
	// environment variable by adding the following line to either your
	// $HOME/.profile or $HOME/.bashrc files:
	// os.Getenv("GREENLIGHT_DB_DSN")
	cfg.db.connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "greenlight")

	db, err := postgres.Connect(cfg.db.connectionString)
	if err != nil {
		log.Fatal(err)
	}

	logger.Println("db connection established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on http://localhost%s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

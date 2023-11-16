package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spargonaut/plant_info_service/internal/data"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config   config
	logger   *log.Logger
	profiles data.Models
}

var validate *validator.Validate

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("PLANTINFO_DB_DSN"), "PostgresSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("database connection pool established")

	app := &application{
		config:   cfg,
		logger:   logger,
		profiles: data.NewModels(db),
	}

	addr := fmt.Sprintf(":%d", cfg.port)
	validate = validator.New(validator.WithRequiredStructEnabled())

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

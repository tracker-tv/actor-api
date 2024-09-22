package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/tracker-tv/actor-api/internal/config"
	"github.com/tracker-tv/actor-api/internal/data"
	"github.com/tracker-tv/actor-api/internal/database"
)

type application struct {
	logger *slog.Logger
	models data.Models
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		logger.Error("failed to parse config", "error", err)
		os.Exit(1)
	}

	logger.Info("starting database connection")

	db, err := database.OpenDB(cfg)
	if err != nil {
		logger.Error("failed to open database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	logger.Info("database connection established")

	app := &application{
		logger: logger,
		models: data.NewModels(db),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/actors", app.GetActors)

	logger.Info("starting server", "port", 8080)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}

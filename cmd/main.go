package main

import (
	"fmt"
	"os"

	"github.com/skurtz97/splat/internal/http"
	"github.com/skurtz97/splat/internal/repo"
	"github.com/skurtz97/splat/internal/splat"
	"go.uber.org/zap"
)

type Config struct {
	Server   http.ServerConfig   `json:"server"`
	Postgres repo.PostgresConfig `json:"postgres"`
}

// run initializes dependencies and runs the program
func run(log *zap.Logger) error {
	log.Info("starting program")

	// database initialization
	db, err := repo.NewPostgresPool()
	if err != nil {
		return fmt.Errorf("failed initializing database: %w", err)
	}
	log.Info("connection to database established")

	// post service initialization
	postService := splat.NewService(log, db)

	// server initialization
	server := http.NewServer(log, postService)
	log.Info("server started on port 8080")

	// start server
	if err := server.Start(); err != nil {
		log.Error("error starting server", zap.Error(err))
		return err
	}

	return nil
}

// main is the entry point for the application, mostly just a wrapper for run
func main() {
	log, _ := zap.NewProduction()

	if err := run(log); err != nil {
		log.Fatal("error running program", zap.Error(err))
		os.Exit(1)
	}
}

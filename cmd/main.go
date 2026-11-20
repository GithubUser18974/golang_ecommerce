package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/sikozonpc/ecom/internal/database"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		slog.Error("DATABASE_URL is required")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	db, err := database.OpenPostgres(ctx, dsn)
	cancel()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	cfg := &config{
		addr: ":8080",
	}
	api := &application{
		config: cfg,
		db:     db,
	}
	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}

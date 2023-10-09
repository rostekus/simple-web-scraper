package main

import (
	"log/slog"
	"os"

	"github.com/rostekus/simple-web-scraper/internal/config"
	"github.com/rostekus/simple-web-scraper/internal/utils/files"
	"github.com/rostekus/simple-web-scraper/internal/utils/logger/sl"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {

	// read config, panic if cannot read
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info(
		"starting web scraper",
		slog.String("env", cfg.Env),
	)
	log.Debug("debug messages are enabled")

	_, err := files.New("urls.txt").Iterator()
	if err != nil {
		log.Error("cannot read file with urls", sl.Err(err))
		os.Exit(1)
	}
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

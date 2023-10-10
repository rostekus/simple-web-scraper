package main

import (
	"log/slog"
	"os"

	"github.com/rostekus/simple-web-scraper/internal/cache"
	"github.com/rostekus/simple-web-scraper/internal/config"
	"github.com/rostekus/simple-web-scraper/internal/controller"
	"github.com/rostekus/simple-web-scraper/internal/scraper"
	"github.com/rostekus/simple-web-scraper/internal/utils/files"
	"github.com/rostekus/simple-web-scraper/internal/utils/logger/sl"
	"github.com/rostekus/simple-web-scraper/internal/words"
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
		slog.Int("maxGo", int(cfg.Scraper.MaxGo)),
		slog.Int("min word lenght", int(cfg.Scraper.MinLen)),
		slog.Int("max word lenght", int(cfg.Scraper.MaxLen)),
	)
	log.Debug("debug messages are enabled")

	urlIter, err := files.New("urls.txt").Iterator()
	if err != nil {
		log.Error("cannot read file with urls", sl.Err(err))
		os.Exit(1)
	}
	outputFilePath := cfg.OutputFilePath

	// setup cache
	inMemoryCache := cache.NewCache()

	resultsChan := make(chan words.WordFreq)
	scraper := scraper.New(log, resultsChan)

	opts := controller.ControllerOpts{
		MaxGo:       int(cfg.MaxGo),
		Cache:       inMemoryCache,
		ResultsChan: resultsChan,
		UrlIter:     urlIter,
		Log:         log,
		Scraper:     scraper,
	}

	c := controller.NewController(&opts)

	// iterate over urls
	// save resaults to resauts chan
	go c.Serve()

	counter := words.NewMostCommonFreqWordsCounter(10) // Replace 10 with your desired number of top words
	for res := range resultsChan {
		counter.Add(res)
	}
	result := counter.Get()
	log.Info("saving file")
	if err := words.SaveWordFreqsToFile(result, outputFilePath); err != nil {
		log.Error("problem with saving file", sl.Err(err))
	}
	log.Info("saved file", slog.String("path", outputFilePath))
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

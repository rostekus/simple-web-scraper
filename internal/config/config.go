package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string `yaml:"env" env-default:"local"`
	OutputFilePath string `yaml:"outputPath" env-default:"output.txt"`
	Scraper        `yaml:"scraper"`
}

type Scraper struct {
	MaxGo  uint `yaml:"maxGo"`
	MaxLen int  `yaml:"maxLen" env-default:"10"`
	MinLen int  `yaml:"minLen" env-default:"1"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

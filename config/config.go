package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

const ApplicationLabel = ""

// Config - app configuration structure.
type Config struct {
	App         string `env:"APP"`
	Environment string `env:"ENVIRONMENT"`
	RedisURL    string `env:"REDIS_URL"`
	CacheTLL    string `env:"CACHE_TLL"`
	HTTPPort    string `env:"HTTP_PORT"`
	HTTPHost    string `env:"HTTP_HOST,default=localhost"`
	MetricsPort string `env:"METRICS_PORT"`
	TracerURL   string `env:"TRACER_URL"`
	LogLevel    string `env:"LOG_LEVEL"`
	BotAPIKey   string `env:"BOT_API_KEY"`
	Services
}

func NewConfig() Config {
	var cfg Config

	if err := envconfig.ProcessWith(context.Background(), &envconfig.Config{
		Target:   &cfg,
		Lookuper: envconfig.OsLookuper(),
	}); err != nil {
		log.Fatal(err)
	}

	return cfg
}

type Services struct {
	Weather    Service `env:"WEATHER_SERVICE_"`
	Translator Service `env:"TRANSLATOR_SERVICE_"`
	News       Service `env:"NEWS_SERVICE_"`
	Currency   Service `env:"CURRENCY_SERVICE_"`
	Definition Service `env:"DEFINITION_SERVICE_"`
}

type Service struct {
	Name    string `env:"PROVIDER"`
	Version string `env:"VERSION"`
	URL     string `env:"URL"`
	APIKey  string `env:"API_KEY"`
}

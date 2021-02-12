package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"unibot-tg/utils"
)

//Config - app configuration structure
type Config struct {
	App string

	LogLevel string

	WeatherKey    string
	TranslatorKey string
	NewsKey       string
	LambasteKey   string
	BotKey        string
	CurrencyKey   string
	DefinitionKey string
}

//Load - load configuration from OS environment
func Load(test bool) *Config {
	if !test {
		err := godotenv.Load()
		utils.FailOnError(err, "Failed to load .env file")
		if err != nil {
			return nil
		}
	} else {
		err := godotenv.Load("../../.env")
		utils.FailOnError(err, "Failed to load .env file")
		if err != nil {
			return nil
		}
	}

	log.Println(".env file was successfully loaded")
	config := Config{}

	config.App = os.Getenv("APP")
	config.BotKey = os.Getenv("BOT_API")
	config.CurrencyKey = os.Getenv("CURRENCY_API")
	config.DefinitionKey = os.Getenv("DEFINITION_API")
	config.LambasteKey = os.Getenv("LAMBASTE_API")
	config.LogLevel = os.Getenv("LOG_LEVEL")
	config.NewsKey = os.Getenv("NEWS_API")
	config.TranslatorKey = os.Getenv("TRANSLATOR_API")
	config.WeatherKey = os.Getenv("WEATHER_API")

	return &config
}

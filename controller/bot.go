package controller

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
	"unibot-tg/config"
	"unibot-tg/services/cli"
	"unibot-tg/services/currency"
	"unibot-tg/services/definition"
	"unibot-tg/services/generator"
	"unibot-tg/services/news"
	"unibot-tg/services/stocks"
	"unibot-tg/services/weather"
	"unibot-tg/utils"
)

type bot struct {
	botObject *tgbotapi.BotAPI
}

var (
	update        = tgbotapi.NewUpdate(0)
	configuration config.Config
)

// NewBot - ...
func NewBot(token string, config config.Config) bot {
	configuration = config

	simpleBot, err := tgbotapi.NewBotAPI(token)
	utils.FailOnError(err, "Failed to create a Bot")

	botStruct := bot{
		simpleBot,
	}

	return botStruct
}

// ConfigureBot - ...
func (b *bot) ConfigureBot(timeout int) {
	if timeout == 0 {
		timeout = 60
	}
	b.botObject.Debug = true

	log.Printf("Authorized on %s", b.botObject.Self.FirstName)

	update.Timeout = timeout
}

//ListenBot - ...
func (b *bot) ListenBot() {
	updates, err := b.botObject.GetUpdatesChan(update)
	utils.FailOnError(err, "Failed to get updates")

	for upd := range updates {
		if upd.Message == nil {
			continue
		} else if !upd.Message.IsCommand() {
			msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)
			_, err = b.botObject.Send(msg)
			utils.FailOnError(err, "Failed to send")
		} else if upd.Message.IsCommand() {
			msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "")
			switch upd.Message.Command() {
			case "help":
				msg.Text = utils.HelpDefaultResponse
			case "start":
				msg.Text = "..."
			case "weather":
				country := strings.Split(upd.Message.CommandArguments(), " ")
				msg.Text = weather.GetWeather(country[0], configuration)
			case "news":
				msg.Text = news.GetNews(upd.Message.CommandArguments(), configuration)
			case "stocks":
				msg.Text = stocks.GetStocks(upd.Message.CommandArguments(), configuration)
			case "currency":
				msg.Text = currency.GetCurrency(upd.Message.CommandArguments(), configuration)
			case "definition":
				msg.Text = definition.GetDefinition(upd.Message.CommandArguments(), configuration)
			case "generator":
				msg.Text = generator.Generate(upd.Message.CommandArguments())
			case "cli":
				msg.Text = cli.Cli(upd.Message.CommandArguments())
			default:
				msg.Text = "Unknown Command"
			}
			if msg.Text == "" {
				msg.Text = "Empty"
			}
			_, errCommand := b.botObject.Send(msg)
			utils.FailOnError(errCommand, "failed to send message")
		}
		log.Printf("[%s] %s", upd.Message.From.UserName, upd.Message.Text)
	}
}

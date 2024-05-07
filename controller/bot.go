package controller

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/s-akhmedoff/unibot-tg/config"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(cfg config.Config) *Bot {
	bot, err := tgbotapi.NewBotAPI(cfg.BotAPIKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Listen() error {
	me, err := b.bot.GetMe()
	if err != nil {
		return err
	}

	log.Println(me.String())

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text+"\nfrom bot")
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := b.bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

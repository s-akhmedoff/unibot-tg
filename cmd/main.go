package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/s-akhmedoff/unibot-tg/config"
	"github.com/s-akhmedoff/unibot-tg/controller"
)

func main() {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	shutdown := make(chan struct{})

	cfg := config.NewConfig()

	go func() {
		Run(cfg)

		<-shutdown

		os.Exit(0)
	}()

	<-shutdownSignal

	close(shutdown)
}

func Run(cfg config.Config) {
	bot := controller.NewBot(cfg)

	go func() {
		if err := bot.Listen(); err != nil {
			log.Fatal(err)
		}
	}()
}

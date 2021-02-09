package main

import (
	"unibot-tg/config"
	"unibot-tg/controller"
)

var configuration config.Config

func init() {
	configInner := config.Load(false)
	if configInner == nil {
		panic("Unable to load configuration")
	}
	configuration = *configInner
}

func main(){
	myBot := controller.NewBot(configuration.BotKey, configuration)
	myBot.ConfigureBot(0)

	myBot.ListenBot()
}

package main

import (
	"log"
	bot "motiv_bot/bot"
	handler "motiv_bot/handlers"
)

func main() {

	token := bot.GetBotToken()
	bot := bot.InitBot(token)

	handler.RegisterHandlers(bot)

	log.Printf("%s started working!\n\n", bot.Me.FirstName)
	bot.Start()
}

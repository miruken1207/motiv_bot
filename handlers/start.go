package handler

import (
	"fmt"
	"log"

	"gopkg.in/telebot.v3"
)

func handleStart(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: %s", c.Sender().FirstName, c.Text())
		
		return c.Send(fmt.Sprintf("Привет, %s 👋!\nДобро пожаловать в %s.", c.Sender().FirstName, bot.Me.FirstName))
	}
}

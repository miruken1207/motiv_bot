package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Проблемы получением токена!")
    }

	token := os.Getenv("BOT_TOKEN")

	settings := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", func(c telebot.Context) error {
		log.Printf("%s: /start\n", c.Sender().FirstName)
		welcome_msg := fmt.Sprintf("Привет, %s 👋!\nДобро пожаловать в %s .",
			c.Sender().FirstName, bot.Me.FirstName)

		return c.Send(welcome_msg)
	})
	
	bot.Handle(telebot.OnText, func(c telebot.Context) error {

		msg := c.Text()
		log.Printf("%s: %s\n", c.Sender().FirstName, msg)
		return c.Send("Вы написали: " + msg)
	})

	log.Printf("%s started working!", bot.Me.FirstName)
	bot.Start()
}
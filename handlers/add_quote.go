package handler

import (
	"log"
	"os"

	"gopkg.in/telebot.v3"
)

var waitingForQuote = make(map[int64]bool)

func appendQuote(path string, quote string) error {

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(quote + "\n")
	return err
}

func handleAddQuote(c telebot.Context) error {

	waitingForQuote[c.Sender().ID] = true

	log.Printf("%s: %s", c.Sender().FirstName, c.Text())

	return c.Send("Отправь мне свою цитату, и я её сохраню ✍️")
}

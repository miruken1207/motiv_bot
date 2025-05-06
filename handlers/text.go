package handler

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func handleText(c telebot.Context) error {

	userID := c.Sender().ID
	text := c.Text()

	if waitingForQuote[userID] {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Ошибка загрузки .env файла: ", err)
		}
		path := os.Getenv("QUOTE_PATH")
		err := appendQuote(path, text)
		if err != nil {
			log.Printf("Ошибка при сохранении цитаты от %s: %v", c.Sender().FirstName, err)
			return c.Send("Произошла ошибка при сохранении цитаты 😞")
		}

		delete(waitingForQuote, userID)
		log.Printf("Пользователь %s добавил цитату: %s", c.Sender().FirstName, text)
		return c.Send("Спасибо! Твоя цитата добавлена 🙏")
	}

	userName := c.Sender().FirstName
	userMessage := c.Text()

	log.Printf("%s: %s", userName, userMessage)
	return c.Send("Вы написали: " + userMessage)
}

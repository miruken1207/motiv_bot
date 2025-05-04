package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

var waitingForQuote = make(map[int64]bool)

func registerHandlers(bot *telebot.Bot) {
	bot.Handle("/start", handleStart(bot))
	bot.Handle(telebot.OnText, handleText)
	bot.Handle("/get_random_quote", handleGetRandomQuote)
	bot.Handle("/add_own_quote", handleAddQuote)
}

func handleStart(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		firstName := c.Sender().FirstName
		log.Printf("%s: /start", firstName)

		welcomeMsg := fmt.Sprintf("Привет, %s 👋!\nДобро пожаловать в %s.", firstName, bot.Me.FirstName)
		return c.Send(welcomeMsg)
	}
}

func handleText(c telebot.Context) error {

	userID := c.Sender().ID
	text := c.Text()

	if waitingForQuote[userID] {
		err := AppendQuote("/home/miruken/Dev/go_dev/motiv_bot/cmd/quotes.txt", text)
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

func GetRandomQuote(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Ошибка открытия файла цитат: %v", err)
		return "Цитаты временно недоступны."
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			quotes = append(quotes, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка чтения файла цитат: %v", err)
		return "Ошибка чтения цитат."
	}

	if len(quotes) == 0 {
		return "Цитаты не найдены."
	}

	rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))]
}

func handleGetRandomQuote(c telebot.Context) error {

	quote := GetRandomQuote("/home/miruken/Dev/go_dev/motiv_bot/cmd/quotes.txt")
	return c.Send(quote)
}

func AppendQuote(path string, quote string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(quote + "\n")
	return err
}

func handleAddQuote(c telebot.Context) error {
	userID := c.Sender().ID
	waitingForQuote[userID] = true

	return c.Send("Отправь мне свою цитату, и я её сохраню ✍️")
}

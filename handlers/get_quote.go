package handler

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func getRandomQuote(path string) string {

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

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла: ", err)
	}

	log.Printf("%s: %s", c.Sender().FirstName, c.Text())

	return c.Send(getRandomQuote(os.Getenv("QUOTE_PATH")))
}

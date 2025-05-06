package handler

import (
	"gopkg.in/telebot.v3"
)

func RegisterHandlers(bot *telebot.Bot) {
	bot.Handle("/start", handleStart(bot))
	bot.Handle(telebot.OnText, handleText)
	bot.Handle("/get_random_quote", handleGetRandomQuote)
	bot.Handle("/add_own_quote", handleAddQuote)
}

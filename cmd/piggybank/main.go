package main

import (
	"log"
	"os"
	"piggybank/internal/bot"
)

func main() {
	token := os.Getenv("PIGGY_BOT_TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("PIGGY_BOT_TELEGRAM_TOKEN не установлен в переменных окружения")
	}

	bot, err := bot.NewBot(token)
	if err != nil {
		log.Fatalf("Ошибка запуска бота: %v", err)
	}

	bot.Start()
}

package main

import (
	"log"
	"os"
	"piggybank/internal/bot"
	"piggybank/internal/db"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

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

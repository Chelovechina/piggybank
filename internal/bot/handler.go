package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	botAPI.Debug = false

	log.Printf("Авторизирован на акк: %s", botAPI.Self.UserName)

	return &Bot{api: botAPI}, nil
}

func (bot *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			bot.handleMessage(update.Message)
		}
	}

}

func (bot *Bot) handleMessage(message *tgbotapi.Message) {
	bot.reply(message.Chat.ID, message.Text)
}

func (bot *Bot) reply(chatID int64, text string) {
	message := tgbotapi.NewMessage(chatID, text)
	_, err := bot.api.Send(message)
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
}

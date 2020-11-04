package cokutbot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CokutBot struct {
	bot *tgbotapi.BotAPI
}

func (c CokutBot) SendMessage(content string) error {
	msg := tgbotapi.NewMessage(-1001485526533, content)
	if _, err := c.bot.Send(msg); err != nil {
		log.Println(err)

		return err
	}

	return nil
}

func NewCokutBot() (CokutBot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("COKUT_BOT_KEY"))
	if err != nil {
		log.Println(err)

		return CokutBot{}, err
	}

	return CokutBot{bot: bot}, nil
}

func RegisterNewBot(botChannel chan string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("COKUT_BOT_KEY"))
	if err != nil {
		log.Println(err)

		return
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	for order := range botChannel {
		msg := tgbotapi.NewMessage(-1001485526533, order)

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)

			continue
		}
	}

	log.Println("BOT END")
}

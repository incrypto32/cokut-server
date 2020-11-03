package cokutbot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RegisterNewBot(botChannel chan string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("COKUT_BOT_KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	for order := range botChannel {
		log.Println("BOT VANNNEEY")

		msg := tgbotapi.NewMessage(-1001485526533, order)

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)

			continue
		}
	}

	log.Println("BOT END")
}

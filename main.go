package main

import (
	"flag"
	"log"
    "telegram_bot_link/consumer/evant-consumer"

	tgClient "telegram_bot_link/clients/telegram"
	
	"telegram_bot_link/events/telegram"
	"telegram_bot_link/storage/files"
)
const (
	tgBotHost = "api.telegram.org"
    storagePath = "storage"
    batchSize = 100
)
func main() {
    eventsProcessor := telegram.New(
        tgClient.New(tgBotHost, mustToken()),
        files.New(storagePath),
    )

    log.Print("Запуск сервиса")

    consumer := evant_consumer.New(eventsProcessor, eventsProcessor, batchSize)

    if err:= consumer.Start();err != nil {
        log.Fatal("сервис остановился")
    }

}


func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}


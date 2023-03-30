package telegram

import (
	"telegram_bot_link/clients/telegram"
)

type Processor struct {
	tg *telegram.Client
	offset int
}

func New (client *telegram.Client, storage)
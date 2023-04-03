package telegram

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strings"

	"telegram_bot_link/lib/e"
	"telegram_bot_link/storage"

)
const (
	RndCmd = "/rnd"
	HelpCmd = "/help"
	StartCmd = "/start"
)
func(p *Processor) doCmd(text string, chatID int,username string)error {
	text= strings.TrimSpace(text)
	log.Printf("got new coomand '%s' from '%s'",text, username )

	if IsAddCmd(text){



	}

	switch text {
	case RndCmd:

	case HelpCmd:

	case StartCmd:

	}
    default:
}

func (p *Processor) savePage(chatID int, pageURl string,  username string ) (err error) {
	defer func() { err = e.WrapIfErr("canr do command:save page ", err) }()

	page := &storage.Page{
		URL: pageURl,
		UserName: username
	}

	IsExists, err := p.

	
}


func IsAddCmd(text string) bool {
	return isURL(text)
	
}

func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
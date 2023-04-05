package telegram

import (
	"context"
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
func(p *Processor) doCmd(text string, chatID int,username string) error {
	text= strings.TrimSpace(text)
	log.Printf("got new coomand '%s' from '%s'",text, username )


	if IsAddCmd(text){
		return p.savePage(chatID,text , username)
	}

	switch text {
	case RndCmd:

	case HelpCmd:

	case StartCmd:

    default:
		return p.tg.SendMassage(chatID, msgUnknownCommand)
	}

	return nil

}

func(p *Processor) savePage(chatID int, pageURl string,  username string ) (err error) {
	defer func() { err = e.WrapIfErr("canr do command:save page ", err) }()

	page := &storage.Page{
		URL:      pageURl,
		UserName: username,
	}
	
	isExists, err := p.storage.IsExists(context.Background(), page)
	if err != nil {
		return err
	}

	if isExists{
		return p.tg.SendMassage(chatID,msgAlreadyExists)
	}

	if err := p.storage.Save(context.Background(), page); err != nil {
		return err
	}

	if err := p.tg.SendMassage(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}


func IsAddCmd(text string) bool {
	return isURL(text)
	
}

func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
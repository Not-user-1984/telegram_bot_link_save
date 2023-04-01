package telegram

import (
	"telegram_bot_link/clients/telegram"
	"telegram_bot_link/events"
	"telegram_bot_link/lib/e"
	"telegram_bot_link/storage"
)

type Processor struct {
	tg *telegram.Client
	offset int
	storage storage.Storage
}

func New(client *telegram.Client, storage storage.Storage) *Processor{
	return &Processor{
		tg: client,
		storage: storage,
	}
}
func(p *Processor) Fetch(limit int)([]events.Event,error) {
	update, err := p.tg.Updates(p.offset,limit)
	if err!= nil {
		return nil, e.Wrap("cant get events", err)
	}
	res := make([]events.Event,0,len(update))

	for _, u := range update{
		res=append(res, event(u))
	}

}

func event(upd telegram.Update) string {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
	if updType==events. {}
	
}

func fetchType(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text

}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return events.Unknow
	}
	return upd.Message.Text
}
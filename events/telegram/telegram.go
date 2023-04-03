package telegram

import (
	"errors"
	"telegram_bot_link/clients/telegram"
	"telegram_bot_link/events"
	"telegram_bot_link/lib/e"
	"telegram_bot_link/storage"
)

type Processor struct {
	tg       *telegram.Client
	offset   int
	storage  storage.Storage
}

type Meta struct {
	offset int 
	storage.Storage
}

var (
	errUnknownEventType = errors.New("unknown event type")
	errUnknownMetaType = errors.New("unknown event type")
)

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

func(p *Processor) ProcessorMessage(event events.Event) {
	meta , err := meta(event)
}




func meta(event events.Event)(Meta, error) {
	res, оk := event.Meta.(Meta)
	if !оk {
		return Meta{}, e.Wrap("cant get meta", errUnknownMetaType)

	}
	return res, nil
}

func event(upd telegram.Update) string {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
	if updType==events.Message{}
	
}

func fetchType(upd telegram.Update) string {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message

}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}
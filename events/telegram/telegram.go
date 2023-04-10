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
	chatID int
	Usermame string
}

var (
	errUnknownEventType = errors.New("unknown event type")
	errUnknownMetaType = errors.New("unknown meta type")
)

func New(client *telegram.Client, storage storage.Storage) *Processor{
	return &Processor{
		tg: client,
		storage: storage,
	}
}


// Данный код является методом Fatch,
// определенным для структуры Processor 
// Метод принимает два аргумента: 
// указатель на объект типа Processor (p) и целочисленное значение limit.
// Метод возвращает список объектов типа events.Event и ошибку.
// Вызывается метод Updates объекта tg (поле структуры Processor),
// чтобы получить обновления из Telegram API.
// Если возникла ошибка при получении обновлений,
// метод возвращает ошибку с помощью функции e.Wrap.
// Если длина списка обновлений равна 0, метод возвращает nil, nil.
// В противном случае создается пустой список res типа events.Event.
// Для каждого обновления в списке обновлений создается объект events.Event
// с помощью функции event и добавляется в список res с помощью функции append.
// Поле offset объекта p устанавливается равным ID последнего обновления + 1.
// Метод возвращает список res и nil в качестве ошибки.
func(p *Processor) Fatch(limit int)([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset,limit)
	if err!= nil {
		return nil, e.Wrap("cant get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event,0,len(updates))

	for _, u := range updates{
		res=append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1
	return res, nil
}

func (p *Processor) Process(event events.Event) error{
	switch event.Type{
	case events.Message:
		return p.ProcessorMessage(event)
	default:
		return e.Wrap("can t process massage", errUnknownEventType)
	}
}

func(p *Processor) ProcessorMessage(event events.Event) error {
	meta , err := meta(event)
	if err != nil {
		return e.Wrap("can t process massege", err)
	}
	

	if err := p.doCmd(event.Text, meta.chatID , meta.Usermame); err != nil {
		return e.Wrap("can t process masage", err)
	}

	return nil

}

func meta(event events.Event)(Meta, error) {
	res, оk := event.Meta.(Meta)
	if !оk {
		return Meta{}, e.Wrap("cant get meta", errUnknownMetaType)

	}
	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
	if updType==events.Message{
		res.Meta = Meta{
			chatID: upd.Message.Chat.ID,
			Usermame: upd.Message.From.Username,
		}
	}
	return res
}

func fetchType(upd telegram.Update) events.Type {
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
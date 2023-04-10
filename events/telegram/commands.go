package telegram

import (
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

// Функция doCmd() обрабатывает команды в чат-приложении,
// вызывая соответствующий метод структуры Processor в зависимости от текста команды. 
// Она принимает указатель на структуру Processor,
// строку с текстом команды и значения chatID и username.
func(p *Processor) doCmd(text string, chatID int,username string) error {
	text= strings.TrimSpace(text)
	log.Printf("got new coomand '%s' from '%s'",text, username )

	if isAddCmd(text){
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID,)
	case StartCmd:
		return p.sendHello(chatID)
    default:
		return p.tg.SendMassage(chatID, msgUnknownCommand)
	}
}

// savePage()Проверяются входные параметры, если какой-либо параметр не верен, то возвращается ошибка.
// Создается объект Page с переданным URL и именем пользователя.
// Проверяется, существует ли уже сохраненная страница в базе данных. Если страница уже сохранена, то возвращается сообщение об ошибке.
// Сохраняется страница в базу данных.
// Отправляется сообщение об успешном сохранении страницы.
// Возвращается nil, если операция выполнена успешно.
func (p *Processor) savePage(chatID int, pageURL string, username string) error {
	if chatID == 0 || pageURL == "" || username == "" {
        return e.New("savePage: invalid input parameter(s)")
    }
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return e.WrapIfErr("can't do command: save page", err)
	}

	if isExists {
		if err := p.tg.SendMassage(chatID, msgAlreadyExists); err != nil {
			return e.WrapIfErr("can't send message: already exists", err)
		}
		return nil
	}

	if err := p.storage.Save(page); err != nil {
		return e.WrapIfErr("can't save page", err)
	}

	if err := p.tg.SendMassage(chatID, msgSaved); err != nil {
		return e.WrapIfErr("can't send message: saved", err)
	}

	return nil
}

// sendRandom() структуры Processor выбирает случайную сохраненную страницу из хранилища
// и отправляет ее URL в указанный чат. Она принимает два параметра: целочисленный chatID и строковый username.
// Функция использует метод PickRandom() структуры storage для выбора случайной страницы из хранилища
// и отправляет ее URL в чат, используя метод SendMassage() структуры tg.
// Если в хранилище нет сохраненных страниц, функция отправляет сообщение "no saved pages" в чат.
// Затем она удаляет выбранную страницу из хранилища с помощью метода Remove() структуры storage.
// Функция возвращает ошибку, если произошла ошибка при выборе случайной страницы или отправке сообщения.
func (p *Processor) sendRandom(chatID int, username string)(err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random",err)}()

	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}

	if errors.Is(err,storage.ErrNoSavedPages) {
		return p.tg.SendMassage(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMassage(chatID,page.URL); err != nil{
		return err
	}
	return p.storage.Remove(page)
}

// sendHelp() структуры Processor отправляет справочную информацию в указанный чат.
func (p *Processor) sendHelp(chatID int) error{
	return p.tg.SendMassage(chatID, msgHelp)
}

// sendHello() структуры Processor отправляет приветственное сообщение в указанный чат
func (p *Processor) sendHello(chatID int) error{
	return p.tg.SendMassage(chatID, msgHello)
}



func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
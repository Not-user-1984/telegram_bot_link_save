package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"telegram_bot_link/lib/e"
)

// Интерфейс Storage содержит четыре метода, которые должны быть реализованы в любом типе,
//  который хочет использовать этот интерфейс. Эти методы позволяют сохранять страницу в хранилище,
//  выбирать случайную страницу для указанного пользователя,
//  удалять страницу из хранилища и проверять, существует ли страница в хранилище.
type Storage interface {
	Save(p *Page)error
	PickRandom(userName string)(*Page, error)
	Remove(p *Page)error
	IsExists(p *Page)(bool, error)
}

var ErrNoSavedPages = errors.New("no saved pages ")

// Структура данных Page представляет веб-страницу и имеет два поля:
// URL, содержащее адрес страницы, и UserName,
// содержащее имя пользователя, который её добавил.
type Page struct {
	URL string
	UserName string
}

// Этот код определяет метод Hash() для структуры данных Page, 
// который вычисляет и возвращает SHA-1 хеш-сумму URL и имени пользователя веб-страницы.
// Сначала создается новый объект sha1 хеш-функции с помощью sha1.New().
// Затем URL и имя пользователя записываются в хеш-функцию с помощью io.WriteString().
// Если при записи происходит ошибка, то вызывается функция e.Wrap() для обработки ошибки и добавления контекстной информации.
// В конце, хеш-сумма рассчитывается с помощью h.Sum(nil)
// и форматируется как шестнадцатеричная строка с помощью fmt.Sprintf("%x",...).

// Метод Hash() возвращает полученную хеш-сумму как строку и nil,
//  если все операции выполняются успешно, или "" и ошибку,
//  если возникает ошибка во время вычисления хеш-суммы.
func (p Page) Hash()(string, error) {
	h := sha1.New()
	
	if _, err := io.WriteString(h,p.URL);err!=nil{
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h,p.UserName);err!=nil{
		return "", e.Wrap("can't calculate hash", err)	
    }
	return fmt.Sprintf("%x",h.Sum(nil)), nil
}
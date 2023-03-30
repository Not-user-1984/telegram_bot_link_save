package files

import (
	"math/rand"
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"
	"telegram_bot_link/lib/e"
	"telegram_bot_link/storage"
	"time"
	"fmt"
)
const (
	defaultPerm = 0774
)
type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}


// Этот код определяет метод Save() для типа Storage,
// который сохраняет веб-страницу, представленную объектом storage.Page,в файловой системе.
// Сначала, определяется путь к файлу, где будет храниться страница.
// Путь формируется с использованием базового пути из поля s.basePath и имени пользователя из поля page.UserName.
// Если директория по этому пути не существует,она создается с помощью os.MkdirAll(). 
// Если возникает ошибка при создании директории, метод Save() завершается с ошибкой.
// Затем, генерируется имя файла с помощью функции fName(), которая вычисляет SHA-1 хеш-сумму URL и имени пользователя page.
// Если возникает ошибка при вызове fName(), метод Save() завершается с ошибкой.
// Далее, полный путь к файлу формируется путем объединения базового пути и имени файла. Файл создается с помощью os.Create().
// Если возникает ошибка при создании файла, метод Save() завершается с ошибкой.
// Теперь, используя gob.NewEncoder() и Encode(), объект page сериализуется в бинарный формат и записывается в файл. 
// Если произошла ошибка при сериализации или записи файла, метод Save() завершается с ошибкой.
// Последнее выражение return nil указывает на успешное завершение метода Save().
// В конце функции используется defer для закрытия файла с помощью file.Close().
// Если произошла ошибка при записи файла, файл в любом случае будет закрыт перед возвратом ошибки.
func(s Storage)  Save(page *storage.Page) (err error) {
	defer func() {err = e.WrapIfErr("cant save page", err)}()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath,defaultPerm);err!=nil{
		return err
	}

	fName,err := fName(page)
	if err!=nil{
		return err
	}

	fPath = filepath.Join(fPath, fName)
	file,err:=os.Create(fPath)
	if err != nil{
		return err
	}

	defer func(){_ =file.Close()}()

	if err := gob.NewEncoder(file).Encode(page);err!=nil{
		return err
	}
	return nil
}

// Этот код возвращает случайную сохраненную страницу из каталога,
// соответствующего заданному пользователю userName.
// Если такие страницы не найдены, то функция возвращает ошибку.
// Функция формирует путь к каталогу, соответствующему заданному пользователю.
// Функция читает список файлов в каталоге и выбирает случайный файл.
// Если список файлов пустой, функция возвращает ошибку.
// В противном случае функция декодирует выбранный файл и возвращает результат как объект *storage.Page.
// Также заметим, что функция использует обработку ошибок с помощью defer
// и вызывает метод WrapIfErr для добавления контекстной информации к ошибке в конце функции перед ее возвратом.
func(s Storage)PickRandom(userName string)(page *storage.Page,err error){
	defer func() {err = e.WrapIfErr("cant pick random", err)}()

	path := filepath.Join(s.basePath, userName)
	files ,err := os.ReadDir(path)
	if err!= nil{
		return nil, err
	}
	if len(files) ==0 {
		return nil, errors.New("no saved page")
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))

}

// Этот код удаляет файл, соответствующий заданному объекту storage.Page
// Функция получает имя файла, связанного с заданным объектом storage.Page, используя метод fName.
// Если возникает ошибка при получении имени файла, функция возвращает ошибку.
// Функция формирует путь к файлу, используя базовый путь и имя пользователя, сохраненные в объекте storage.Page,
// и имя файла, полученное на предыдущем шаге.
// Функция удаляет файл по указанному пути. Если происходит ошибка удаления файла,
// функция оборачивает ее в контекстную информацию и возвращает как ошибку.
// Если удаление прошло успешно, функция возвращает nil.
func (s Storage) Remove(p *storage.Page) error{
	fileName,err := fName(p)
	if err != nil {
		return e.Wrap("can't remove file", err)
	}
	path := filepath.Join(s.basePath, p.UserName,fileName)

	if err := os.Remove(path);err != nil{
		msg := fmt.Sprintf("can't remove file %s", path)
		return e.Wrap(msg,err)
	}
	return nil
}


// Эта функция принимает на вход объект Storage и указатель на объект типа storage.Page.
// Она проверяет, существует ли файл, соответствующий переданному объекту storage.Page.
// Сначала функция вызывает вспомогательную функцию fName, которая получает имя файла из storage.Page. 
// Если происходит ошибка при получении имени файла, то функция возвращает false и ошибку.
// Затем функция использует filepath.Join для создания пути к файлу,
// используя базовый путь хранилища,имя пользователя и имя файла.
// Далее функция использует оператор switch для проверки наличия файла по указанному пути.
// Если файла не существует, функция возвращает false без ошибки.
// Если произошла другая ошибка, функция возвращает false вместе с ошибкой,
// содержащей сообщение об ошибке и исходную ошибку.
// В конце функция возвращает true, если файл существует, или false, если он отсутствует, а также ошибку (если есть).
func (s Storage)IsExists(p *storage.Page)(bool,error){
	fileName,err := fName(p)
	if err != nil {
		return false,e.Wrap("can't check if file exists",err)
	}
	path := filepath.Join(s.basePath, p.UserName,fileName)
	switch _, err = os.Stat(path); {
	case errors.Is(err,os.ErrNotExist):
		return false, nil

	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exists",path)
		return false, e.Wrap(msg, err)
	}

	return true, nil

}
// Этот код декодирует сохраненный файл страницы по заданному пути filePath
// и возвращает десериализованный объект storage.Page.
// Функция открывает файл по указанному пути filePath с помощью функции os.Open.
// Если возникает ошибка при открытии файла,
// функция оборачивает ее в контекстную информацию с помощью метода e.Wrap и возвращает как ошибку.
// Функция регистрирует функцию закрытия файла с помощью оператора defer,
// чтобы гарантировать, что файл будет закрыт после чтения.
// Декодирование файла происходит с помощью gob.NewDecoder(f).Decode(&p), 
// где f - открытый файл, а &p - адрес объекта storage.Page, в который записывается результат декодирования.
// Если возникает ошибка при декодировании файла,
// функция оборачивает ее в контекстную информацию с помощью метода e.Wrap и возвращает как ошибку.
// Функция возвращает десериализованный объект storage.Page через указатель на него вместе с nil.
func(s Storage) decodePage(filePath string)(*storage.Page, error){
	f, err := os.Open(filePath)

	if err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	defer func(){_ =f.Close()}()

	var p storage.Page
	if err :=gob.NewDecoder(f).Decode(&p);err != nil{
		return nil, e.Wrap("can't decode page",err)
	}
    return &p,nil 
}

func fName(p *storage.Page) (string, error) {
	return p.Hash()
}
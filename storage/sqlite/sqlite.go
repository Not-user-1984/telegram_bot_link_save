package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"telegram_bot_link/storage"
)


type Storage struct {
	db *sql.DB
}


// Метод Save структуры Storage сохраняет новую страницу в таблице pages базы данных.
// Он принимает контекст и указатель на структуру Page в качестве параметров.
// URL и имя пользователя используются в SQL-запросе для проверки,
// существует ли страница уже в базе данных.
// Если запрос успешно выполняется, то страница сохраняется в базе данных и функция возвращает nil.
// Если происходит ошибка, то функция возвращает ошибку с соответствующим сообщением.
func New (path string) (*Storage , error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil , fmt.Errorf("can't open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cant't connect to database: %w" , err )
	}
	return &Storage{db: db} , nil
}
func (s *Storage) Save (ctx context.Context, p *storage.Page) error{
	q := `SELECT COUNT(*) FROM  pages WHERE url = ? AND  user_name = ?`

	if _, err := s.db.ExecContext(ctx, q , p.URL , p.UserName ,); err != nil {
		return fmt.Errorf("can`t save page: %w", err)
	}

	return nil
}


// Метод PickPandom структуры Storage выбирает случайную страницу из таблицы pages для определенного пользователя.
// Он принимает контекст и имя пользователя в качестве параметров.
// В SQL-запросе используются имя пользователя и функция RANDOM(),
// которая выбирает случайный URL из таблицы.
// Если запрос выполняется успешно, URL сохраняется в переменной url.
// Если в таблице нет записей для данного пользователя, метод возвращает ошибку ErrNoSavedPages.
// Если происходит другая ошибка, метод возвращает ошибку с соответствующим сообщением.
// В результате метод возвращает указатель на структуру Page, содержащую URL и имя пользователя случайной страницы.
func (s *Storage) PickPandom (ctx context.Context, userName string) (*storage.Page, error){
	q := `SElECT url FROM pages WHERE user_name = ? ORDER BY RANDOM() LIMIT 1`

	var url string


	err := s.db.QueryRowContext(ctx, q , userName).Scan(&url)
	// context в Go является стандартным пакетом для передачи метаданных запроса через границы API и горoutines.
	// Контекст может использоваться для управления временем жизни запроса,
	// передачи значений/данных между функциями, отмены операции, обработки ошибок и т.д.
	// В данном коде контекст используется для выполнения запроса к базе данных с помощью метода QueryRowContext. Контекст передается этому методу в качестве первого аргумента, он позволяет контролировать время ожидания запроса и отменять его при необходимости.
	// Для создания контекста можно использовать функцию context.Background(),
	// которая возвращает пустой контекст.
	// После этого можно добавлять значения и параметры к контексту с помощью функций-оберток,
	// таких как context.WithCancel() или context.WithTimeout().
	if err == sql.ErrNoRows{
		return nil , storage.ErrNoSavedPages
	}

	if  err != nil {
		return nil, fmt.Errorf("can`t pick random page: %w", err)
	}

	return &storage.Page{
		URL: url,
		UserName: userName,
	}, nil

	}
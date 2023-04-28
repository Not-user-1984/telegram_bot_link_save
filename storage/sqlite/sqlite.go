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

func (s *Storage) Save (ctx context.Context, p *storage.Page)
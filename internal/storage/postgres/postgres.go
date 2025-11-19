package postgres

import (
	"database/sql"
	"fmt"
	 _ "github.com/lib/pq"

)

type Storage struct {
	db *sql.DB
}

func New(host string, port int, user string, password string, dbname string) (*Storage, error) {
	const op = "storage.postgres.New"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return &Storage{}, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return &Storage{}, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println("Successfully connected!")

	return &Storage{db: db}, nil
}

func CreateTable(s *Storage) error {
	const op = "storage.postgres.CreateTable"

	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS numbers (value INT)`)
    if err != nil {
		return fmt.Errorf("%s: %w", op, err)
    }
	return nil
}
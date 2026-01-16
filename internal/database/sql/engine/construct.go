package engine

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabaseEngine(connStr string) (*DatabaseEngine, error) {
	// db, err := sql.Open("sqlite3", fileName)

	// if err != nil {
	// 	return nil, err
	// }
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("База недоступна:", err)
	}

	fmt.Println("Подключение к PostgreSQL успешно!")
	return &DatabaseEngine{
		db: db,
	}, nil
}

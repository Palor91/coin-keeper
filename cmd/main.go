package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("База недоступна:", err)
	}

	fmt.Println("Подключение к PostgreSQL успешно!")
}

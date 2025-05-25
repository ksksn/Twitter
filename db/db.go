package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" 
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "file:twitter.db")
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Println("Соединение с базой данных успешно установлено")

	if err := InitSchema(); err != nil {
		return err
	}
	log.Println("Схема базы данных успешно инициализирована")

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

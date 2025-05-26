package db

import (
	"fmt"
	"log"
)

func InitSchema() error {
	log.Println("Инициализация схемы базы данных...")

	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	log.Println("Таблица users проверена")

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			is_active BOOL DEFAULT 1,
			text TEXT NOT NULL,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	log.Println("Таблица notes проверена")

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS likes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			note_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(note_id) REFERENCES notes(id),
			UNIQUE(user_id, note_id)
		)`)
	if err != nil {
		return err
	}
	log.Println("Таблица likes проверена")

	return nil
}

func LikeNote(userID, noteID int) error {
	// Проверяем, что заметка существует и активна
	var exists bool
	err := DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM notes 
			WHERE id = ? AND is_active = 1
		)
	`, noteID).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("заметка не найдена или удалена")
	}

	// Добавляем лайк (UNIQUE constraint предотвратит дублирование)
	_, err = DB.Exec(`
		INSERT INTO likes (user_id, note_id) 
		VALUES (?, ?)
	`, userID, noteID)

	return err
}

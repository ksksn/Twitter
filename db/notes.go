package db

import (
	"database/sql"
	"twitter/models"
)

func CreateNotes(db *sql.DB, note *models.Note, userID uint) error {
	query := "INSERT INTO notes (text, user_id) VALUES (?, ?)"
	result, err := db.Exec(query, note.Text, userID)
	if err != nil {
		return err
	}

	// Получаем ID созданной записи
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

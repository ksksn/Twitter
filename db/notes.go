package db

import (
	"database/sql"
	"fmt"
	"twitter/models"
)

func CreateNotes(db *sql.DB, note *models.Note, userID uint) error {
	query := "INSERT INTO notes (text, user_id, is_active) VALUES (?, ?, 1)"
	result, err := db.Exec(query, note.Text, userID)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	note.ID = int(lastID)
	note.IsActive = true
	note.UserID = userID
fmt.Println("Note created with ID:", userID)
	return nil
}



func DeleteNotes(db *sql.DB, noteID int, userID uint) error {
	query := "UPDATE notes SET is_active = 0 WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, noteID, userID)
	return err
}

func GetNoteByID(db *sql.DB, noteID int) (*models.Note, error) {
	var note models.Note
	query := "SELECT id, text, is_active, user_id FROM notes WHERE id = ? AND is_active = 1"
	err := db.QueryRow(query, noteID).Scan(&note.ID, &note.Text, &note.IsActive, &note.UserID)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func LikeNotes(db *sql.DB, note *models.Note, userID uint) error {
	query := "INSERT INTO likes (user_id, note_id) VALUES (?, ?)"
	result,err:=db.Exec(query, userID, note.ID)
	if err!= nil {
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
}
note.ID = int(lastID)
note.UserID = userID

return nil
}


func UnlikeNotes(db *sql.DB, note *models.Note, userID uint) error{
	query := "DELETE FROM likes (user_id, note_id) VALUES (?, ?)"
	_,err:=db.Exec(query, userID, note.ID)
	if err!= nil {
		return err
	}
	

	return nil
	}
package db

import (
	"fmt"
	"log"
	"twitter/models"
)

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	query := `SELECT id, username, password FROM users WHERE username = ?`
	err := DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(user *models.User) (uint, error) {
	if user == nil {
		return 0, fmt.Errorf("user object is nil")
	}

	if user.Username == "" {
		return 0, fmt.Errorf("username cannot be empty")
	}

	log.Printf("Создание пользователя с именем: %s", user.Username)

	
	query := `INSERT INTO users (username, password) VALUES (?, ?)`

	result, err := DB.Exec(query, user.Username, user.Password)
	if err != nil {
		return 0, fmt.Errorf("database error on insert: %w", err)
	}


	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error retrieving last insert id: %w", err)
	}

	user.ID = uint(lastID)
	log.Printf("Создан пользователь с ID: %d", lastID)
	return user.ID, nil
}

func UserExists(username string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)`
	err := DB.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

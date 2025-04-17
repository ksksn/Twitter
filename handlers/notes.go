package handlers

import (
    "net/http"
    "twitter/models" // Убедитесь, что этот путь правильный

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
  
)

// Временное хранилище заметок (ключ - ID, значение - текст)
var notesStore = make(map[string]string)

func CreateNote(c *gin.Context) {
    var newNoteRequest models.CreateNote // Используем структуру из models/notes.go

    // Пытаемся привязать JSON из тела запроса к структуре
    if err := c.BindJSON(&newNoteRequest); err != nil {
        // Если есть ошибка валидации (например, отсутствует поле "save"), возвращаем 400

        c.JSON(http.StatusBadRequest, err)
        return
    }

    // Генерируем уникальный ID для заметки
    id := uuid.New().String()

    // Сохраняем текст заметки в наше временное хранилище
    notesStore[id] = newNoteRequest.Text

   

    // Возвращаем JSON с ID и текстом заметки
    c.JSON(http.StatusOK, gin.H{
        "id":   id,
        "note": newNoteRequest.Text, // Возвращаем текст из запроса
    })
}
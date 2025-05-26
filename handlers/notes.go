package handlers

import (
	"log"
	"net/http"
	"strconv"
	"twitter/db"
	"twitter/models"

	"github.com/gin-gonic/gin"
)

var notesStore = make(map[string]string)
var likedNotes = make(map[string]bool) //ааааааа дурацкая ерунда

func CreateNote(c *gin.Context) {
	type request struct {
		Text string `json:"text" binding:"required"`
	}

	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Получаем userID из контекста (устанавливается middleware)
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID format",
		})
		return
	}

	note := &models.Note{
		Text: req.Text,
	}

	if err := db.CreateNotes(db.DB, note, userID); err != nil {
		log.Println("Error creating note:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Ошибка при сохранении заметки",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Заметка успешно сохранена",
	})

}

func DeleteNote(c *gin.Context) {
	type request struct {
		ID string `json:"note_id" binding:"required"`
	}
	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID format",
		})
		return
	}

	// Конвертируем ID заметки в int
	noteID, err := strconv.Atoi(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid note ID",
		})
		return
	}

	// Проверяем, существует ли заметка и принадлежит ли она пользователю
	note, err := db.GetNoteByID(db.DB, noteID)
	if err != nil {
		log.Printf("Error getting note: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}

	// Проверяем, является ли пользователь владельцем заметки
	if note.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You can only delete your own notes",
		})
		return
	}

	// Выполняем мягкое удаление
	if err := db.DeleteNotes(db.DB, noteID, userID); err != nil {
		log.Printf("Error soft deleting note: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
	log.Println("notesStore", notesStore)
}

// func LikeNote(c *gin.Context) {
// 	log.Println("likedNotes", likedNotes)
// 	type request struct {
// 		ID string `json:"note_id" binding:"required"`
// 	}
// 	var req request
// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid request",
// 		})
// 		return
// 	}
// 	_, exists := notesStore[req.ID]
// 	if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Note not found",
// 		})
// 		return
// 	}

// 	_, exists = likedNotes[req.ID]
// 	if exists {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Note have already liked",
// 		})
// 		return
// 	} else {

// 		likedNotes[req.ID] = true
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Note was liked",
// 	})
// 	log.Println("likedNotes", likedNotes)
// }

func DislikeNote(c *gin.Context) {
	log.Println("likedNotes", likedNotes)
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	_, exists := likedNotes[req.ID]
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"message": "Note итак не liked",
		})
		return
	} else {
		delete(likedNotes, req.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note was disliked",
	})
	log.Println("likedNotes", likedNotes)
}

func AboutMe(c *gin.Context) {

}

func LikeNote(c *gin.Context) {
	type request struct {
		ID int `json:"note_id" binding:"required"`
	}

	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Получаем userID из контекста (устанавливается middleware)
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID format",
		})
		return
	}

	note := &models.Note{
		ID: req.ID,
	}
	

	if err := db.LikeNotes(db.DB, note,userID); err != nil {
		log.Println("Error creating note:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Ошибка при сохранении заметки",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Заметка успешно сохранена",
	})

}

func UnlikeNote(c *gin.Context){
	
	type request struct {
        ID int `json:"note_id" binding:"required"`
    }

}
package handlers

import (
	"log"
	"net/http"
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
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	_, exists := notesStore[req.ID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Note not found",
		})
		return
	}

	delete(notesStore, req.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
	log.Println("notesStore", notesStore)
}

func LikeNote(c *gin.Context) {
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
	_, exists := notesStore[req.ID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Note not found",
		})
		return
	}

	_, exists = likedNotes[req.ID]
	if exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Note have already liked",
		})
		return
	} else {

		likedNotes[req.ID] = true
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note was liked",
	})
	log.Println("likedNotes", likedNotes)
}

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

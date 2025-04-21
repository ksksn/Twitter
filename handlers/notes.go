package handlers

import (
    "net/http"
    "twitter/models" 
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "log"
)

var notesStore = make(map[string]string)
var likedNotes = make(map[string]bool) //ааааааа дурацкая ерунда 


func CreateNote(c *gin.Context) {
    
    var newNote models.Note 
    if err := c.BindJSON(&newNote); err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    id := uuid.New().String()
    notesStore[id] = newNote.Text
    log.Println("notesStore", notesStore)
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
    }else{

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
    }else{
        delete(likedNotes, req.ID)
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Note was disliked",
    })
    log.Println("likedNotes", likedNotes)
}

func AboutMe(c *gin.Context){
    
}


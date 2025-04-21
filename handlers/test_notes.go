package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "twitter/models"
    "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/notes", CreateNote)
    r.DELETE("/notes", DeleteNote)
    r.POST("/notes/like", LikeNote)
    r.POST("/notes/dislike", DislikeNote)
    return r
}

func TestCreateNote(t *testing.T) {
    router := setupRouter()

    newNote := models.Note{Text: "This is a test note"}
    jsonValue, _ := json.Marshal(newNote)

    req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonValue))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
    }

    if len(notesStore) == 0 {
        t.Error("Expected notesStore to have 1 note, but got 0")
    }
}

func TestDeleteNote(t *testing.T) {
    router := setupRouter()

    // Создаём заметку
    newNote := models.Note{Text: "Note to delete"}
    jsonValue, _ := json.Marshal(newNote)
    req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonValue))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    id := "" 

    deleteReq, _ := json.Marshal(map[string]string{"id": id})
    req, _ = http.NewRequest("DELETE", "/notes", bytes.NewBuffer(deleteReq))
    response = httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
    }

    if _, exists := notesStore[id]; exists {
        t.Error("Expected note to be deleted, but it still exists")
    }
}

func TestLikeNote(t *testing.T) {
    router := setupRouter()

    newNote := models.Note{Text: "Note to like"}
    jsonValue, _ := json.Marshal(newNote)
    req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonValue))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    id := "" 

    likeReq, _ := json.Marshal(map[string]string{"id": id})
    req, _ = http.NewRequest("POST", "/notes/like", bytes.NewBuffer(likeReq))
    response = httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
    }

    if !likedNotes[id] {
        t.Error("Expected note to be liked, but it is not")
    }
}

func TestLikeNoteAlreadyLiked(t *testing.T) {
    router := setupRouter()

    newNote := models.Note{Text: "Note to like"}
    jsonValue, _ := json.Marshal(newNote)
    req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonValue))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    id := "" 

    likeReq, _ := json.Marshal(map[string]string{"id": id})
    req, _ = http.NewRequest("POST", "/notes/like", bytes.NewBuffer(likeReq))
    response = httptest.NewRecorder()
    router.ServeHTTP(response, req)

    req, _ = http.NewRequest("POST", "/notes/like", bytes.NewBuffer(likeReq))
    response = httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusNotFound {
        t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, status)
    }
}

func TestDislikeNote(t *testing.T) {
    router := setupRouter()

    newNote := models.Note{Text: "Note to dislike"}
    jsonValue, _ := json.Marshal(newNote)
    req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonValue))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    id := "" 

    likeReq, _ := json.Marshal(map[string]string{"id": id})
    req, _ = http.NewRequest("POST", "/notes/like", bytes.NewBuffer(likeReq))
    response = httptest.NewRecorder()
    router.ServeHTTP(response, req)

    dislikeReq, _ := json.Marshal(map[string]string{"id": id})
    req, _ = http.NewRequest("POST", "/notes/dislike", bytes.NewBuffer(dislikeReq))
    response = httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
    }

    if _, exists := likedNotes[id]; exists {
        t.Error("Expected note to be disliked, but it still exists in likedNotes")
    }
}

func TestDislikeNoteNotLiked(t *testing.T) {
    router := setupRouter()

    newNote := models.Note{Text: "Note not liked"}
    jsonValue, _ := json.Marshal(newNote)
    req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonValue))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    id := "" 
    dislikeReq, _ := json.Marshal(map[string]string{"id": id})
    req, _ = http.NewRequest("POST", "/notes/dislike", bytes.NewBuffer(dislikeReq))
    response = httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
    }
}

func TestDeleteNonExistentNote(t *testing.T) {
    router := setupRouter()
    deleteReq, _ := json.Marshal(map[string]string{"id": "non-existent-id"})
    req, _ := http.NewRequest("DELETE", "/notes", bytes.NewBuffer(deleteReq))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusNotFound {
        t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, status)
    }
}

func TestLikeNonExistentNote(t *testing.T) {
    router := setupRouter()

    likeReq, _ := json.Marshal(map[string]string{"id": "non-existent-id"})
    req, _ := http.NewRequest("POST", "/notes/like", bytes.NewBuffer(likeReq))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusNotFound {
        t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, status)
    }
}

func TestInvalidCreateNote(t *testing.T) {
    router := setupRouter()
    newNote := models.Note{Text: ""}
    jsonValue, _ := json.Marshal(newNote)

    req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonValue))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, req)

    if status := response.Code; status != http.StatusBadRequest {
        t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, status)
    }
}
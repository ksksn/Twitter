package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"twitter/models"
	"twitter/utils"
)

var users = make(map[string]models.User)
var userCounter uint = 1
func Register(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := users[input.Username]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user := models.User{
		ID:       userCounter,
		Username: input.Username,
		Password: string(hashedPassword),
	}

	users[input.Username] = user
	userCounter++

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, exists := users[input.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{
		AccessToken: token,
	})
}


func GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}

	var foundUser models.User
	found := false

	for _, user := range users {
		if user.ID == userID.(uint) {
			foundUser = user
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Возвращаем данные пользователя (исключая пароль)
	c.JSON(http.StatusOK, gin.H{
		"id":       foundUser.ID,
		"username": foundUser.Username,
		"about_me": foundUser.AboutMe,
	})
}

// UpdateAboutMe обновляет информацию "обо мне" для аутентифицированного пользователя
func UpdateAboutMe(c *gin.Context) {
	// Получаем ID пользователя из контекста
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}

	// Структура для входных данных
	var input struct {
		AboutMe string `json:"about_me" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновляем информацию пользователя
	for username, user := range users {
		if user.ID == userID.(uint) {
			user.AboutMe = input.AboutMe
			users[username] = user
			c.JSON(http.StatusOK, gin.H{"message": "about me updated successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

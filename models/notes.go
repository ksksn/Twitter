package models

type Note struct {
	ID       int    `json:"id"`
	Text     string `json:"text" binding:"required"`
	IsActive bool   `json:"is_active"`
	UserID   uint   `json:"user_id"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Like struct {
	ID     int  `json:"id"`
	UserID uint `json:"user_id"`
	NoteID int  `json:"note_id"`
}

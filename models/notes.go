package models

type Note struct {
	ID   int `json:"id"`
	Text string `json:"text" binding:"required"`
	//IsLiked bool   `json:"is_liked"`
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

package models

type Note struct {
	Text string `json:"text" binding:"required"`
	//IsLiked bool   `json:"is_liked"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	AboutMe  string `json:"about_me"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

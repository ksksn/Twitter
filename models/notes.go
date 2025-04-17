package models

type CreateNote struct {
    Text string `json:"text" binding:"required"`
}

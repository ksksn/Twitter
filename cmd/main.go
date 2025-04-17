package main

import (
    "github.com/gin-gonic/gin"
  
    "twitter/handlers"
    
)

func main() {
    router := gin.Default()

    router.POST("/create", handlers.CreateNote) 

    router.Run(":8080") 
}


//удаление структурку с конфигом, структуру  с роутом,где я буду указывать все ендпоинты
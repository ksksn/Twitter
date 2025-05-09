package routes

import (
	"twitter/handlers"
	"twitter/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}

func (r *Router) SetupRoutes() {
	auth := r.engine.Group("api/v1/auth")
	{
		auth.POST("register", handlers.Register)
		auth.POST("login", handlers.Login)
	}

	// Защищенные маршруты для заметок
	notes := r.engine.Group("api/v1/notes")
	notes.Use(middleware.AuthMiddleware())
	{
		notes.POST("create", handlers.CreateNote)
		notes.POST("like", handlers.LikeNote)
		notes.POST("dislike", handlers.DislikeNote)
		notes.DELETE("delete", handlers.DeleteNote)
	}

	users := r.engine.Group("api/v1/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("profile", handlers.GetProfile)
		users.POST("aboutme", handlers.UpdateAboutMe)
	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}


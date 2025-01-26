package routes

import (
	handlers "blogging-platform/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	g := gin.Default()
	api := g.Group("/api")
	posts := api.Group("/posts")

	postsHandler := handlers.NewPostHandler()
	posts.GET("", postsHandler.GetPosts)
	posts.GET("/:id", postsHandler.GetPostByID)
	posts.POST("", postsHandler.CreatePost)
	posts.PUT("/:id", postsHandler.UpdatePost)
	posts.DELETE("/:id", postsHandler.DeletePost)
	g.Run()
}
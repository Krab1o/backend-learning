package server

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	api := r.Group("/api")
	static := r.Group("/")
	static.Static("/static", "./static")
	r.LoadHTMLGlob("static/templates/*")
	
	api.POST("/login", (&LoginHandler{}).Login)
	api.GET("/articles", (&ArticlesHandler{}).GetArticle)
	api.GET("/articles/:id", (&ArticlesHandler{}).GetArticleByID)
	api.POST("/articles", (&ArticlesHandler{}).CreateArticle)
	api.PUT("/articles/:id", (&ArticlesHandler{}).UpdateArticle)
	api.DELETE("/articles/:id", (&ArticlesHandler{}).DeleteArticle)
	
	static.GET("/home", (&HomeHandler{}).PageHome)
	static.GET("/login", (&LoginHandler{}).PageLogin)
	static.GET("/articles", (&ArticlesHandler{}).PageArticles)
}
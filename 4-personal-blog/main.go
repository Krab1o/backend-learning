package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Login 		string	`json:"login"`
	Password 	string	`json:"password"`
} 

type LoginHandler struct {}

func (l *LoginHandler) Login(c *gin.Context) {
	
	c.String(http.StatusOK, "Wrong Password")
}

func (l *LoginHandler) GetPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Init(r *gin.Engine) {
	loginHandler := LoginHandler{}
	//TODO: group methods
	api := r.Group("/api")
	static := r.Group("/")
	static.Static("/static", "./static")
	r.LoadHTMLGlob("static/templates/*")
	
	api.POST("/login", loginHandler.Login)
	static.GET("/login", loginHandler.GetPage)
}

func main() {
	r := gin.Default()
	Init(r)
	r.Run()
}
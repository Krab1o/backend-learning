package server

import (
	"encoding/json"
	"log"
	"net/http"
	"personal-blog/internal/db"
	"personal-blog/internal/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {}
type ArticlesHandler struct {}
type LoginHandler struct {}

func (h *HomeHandler) PageHome(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "http://localhost:8080/articles")
}

func (l *LoginHandler) Login(c *gin.Context) {
	user := &types.User{}
	json.NewDecoder(c.Request.Body).Decode(user)
	if (user.Email == "example@gmail.com" && 
		user.Password == "admin") {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong password",
		})
	}
}

func (l *LoginHandler) PageLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (a *ArticlesHandler) PageArticles(c *gin.Context) {
	c.HTML(http.StatusOK, "articles.html", nil)
}

func (a *ArticlesHandler) GetArticle(c *gin.Context) {
	log.Println(c.Request.RequestURI)

	c.JSON(http.StatusOK, db.GetArticles())
}

func (a *ArticlesHandler) GetArticleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, ok := db.GetArticleByID(uint(id))
	if ok {
		c.JSON(http.StatusOK, article)
	} else {
		c.Status(http.StatusNotFound)
	}
	
	log.Println(c.Request.RequestURI)
}

func (a *ArticlesHandler) CreateArticle(c *gin.Context) {
	newArticle := types.Article{}
	json.NewDecoder(c.Request.Body).Decode(&newArticle)
	db.CreateArticle(&newArticle)

	c.Status(http.StatusCreated)
	log.Println(c.Request.RequestURI)
}

func (a *ArticlesHandler) UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	newArticle := types.Article{}
	json.NewDecoder(c.Request.Body).Decode(&newArticle)
	ok := db.UpdateArticle(uint(id), &newArticle)
	if ok {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusNotFound)
	}
	log.Println(c.Request.RequestURI)
}

func (a *ArticlesHandler) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ok := db.DeleteArticle(uint(id))
	if ok {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusNotFound)
	}
	
	log.Println(c.Request.RequestURI)
}
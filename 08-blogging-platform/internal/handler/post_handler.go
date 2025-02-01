package handler

import (
	"blogging-platform/internal/model"
	"blogging-platform/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler() *postHandler{
	service := service.NewPostService()
	return &postHandler{service: service}
}

// Log are redundant but written on educational purposes to follow separation
// of three layers: handler, service and repository

func (p *postHandler) GetPosts(c *gin.Context) {
	posts, err := p.service.GetPosts()
	if err != nil {
		log.Printf("Handler get posts: failed to get posts, %v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, posts)
}

func (p *postHandler) GetPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Handler get one post: failed to convert ID, %v", err)
		c.Status(http.StatusNotFound)
	}
	post, err := p.service.GetPostByID(uint(id))
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, post)
}

func (p *postHandler) CreatePost(c *gin.Context) {
	newPost := &model.Post{}
	err := c.ShouldBindJSON(newPost)
	if err != nil {
		log.Printf("Handler create post: failed to convert body, %v", err)
		c.Status(http.StatusBadRequest)
	}
	err = p.service.CreatePost(newPost)
	if err != nil {
		log.Printf("Handler create post: failed to create post, %v", err)
	}
	c.Status(http.StatusCreated)
}

func (p *postHandler) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := p.service.DeletePost(uint(id))
	if err != nil {
		c.Status(http.StatusNotFound)	
		return
	}
	c.Status(http.StatusNoContent)
}

func (p *postHandler) UpdatePost(c *gin.Context) {
	newPost := model.Post{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Handler update post: failed to convert ID, %v", err)
		c.Status(http.StatusNotFound)
	}
	err = c.ShouldBindJSON(&newPost)
	if err != nil {
		log.Printf("Handler update post: failed to convert body, %v", err)
		c.Status(http.StatusBadRequest)
	}
	newPost.ID = uint(id)
	err = p.service.UpdatePost(&newPost)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}


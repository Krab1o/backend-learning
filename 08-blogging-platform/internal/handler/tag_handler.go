package handler

import (
	"blogging-platform/internal/model"
	"blogging-platform/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type tagHandler struct {
	service service.TagService
}

func NewTagHandler() *tagHandler{
	service := service.NewTagService()
	return &tagHandler{service: service}
}

func (p *postHandler) GetTags(c *gin.Context) {
	posts, err := p.service.GetPosts()
	if err != nil {
		log.Printf("Handler: failed to get posts, %v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, posts)
}

func (p *postHandler) GetTagByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Handler: failed to convert ID, %v", err)
	}
	post, err := p.service.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, post)
		return
	}
	c.JSON(http.StatusOK, post)
}

func (p *postHandler) CreatePost(c *gin.Context) {
	newPost := &model.Post{}
	err := c.ShouldBindJSON(newPost)
	if err != nil {
		log.Printf("Handler: failed to convert body, %v", err)
	}
	err = p.service.CreatePost(newPost)
	if err != nil {
		log.Printf("Handler: failed to create post, %v", err)
	}
	c.Status(http.StatusCreated)
}

func (p *postHandler) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := p.service.DeletePost(uint(id))
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}
	c.Status(http.StatusNotFound)
}

func (p *postHandler) UpdatePost(c *gin.Context) {
	newPost := model.Post{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Handler: failed to convert ID, %v", err)
	}
	err = c.ShouldBindJSON(&newPost)
	if err != nil {
		log.Printf("Handler: failed to convert body, %v", err)
	}
	err = p.service.UpdatePost(uint(id), &newPost)
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}
	c.Status(http.StatusNotFound)
}

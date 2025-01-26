package handlers

import (
	"blogging-platform/internal/models"
	"blogging-platform/internal/stores"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	store stores.PostStorage
}

func NewPostHandler() *postHandler{
	store := stores.NewPostStorageDB()
	return &postHandler{store: store}
}

// For now I'll write all logic in handlers
// TODO: Refactor code and move it to "service" and "repository" folders

func (p *postHandler) GetPosts(c *gin.Context) {
	posts := p.store.Get()
	c.JSON(http.StatusOK, posts)
}

func (p *postHandler) GetPostByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, ok := p.store.GetByID(uint(id))
	if ok {
		c.JSON(http.StatusOK, post)
	} else {
		c.JSON(http.StatusNotFound, post)
	}
}

func (p *postHandler) CreatePost(c *gin.Context) {
	newPost := models.Post{}
	err := c.ShouldBindJSON(&newPost)
	if (err != nil) {
		log.Println(err)
	}
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = newPost.CreatedAt	

	p.store.Add(&newPost)
	c.Status(http.StatusCreated)
}

func (p *postHandler) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ok := p.store.Delete(uint(id))
	if ok {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (p *postHandler) UpdatePost(c *gin.Context) {
	newPost := models.Post{}
	err := c.ShouldBindJSON(&newPost)
	if (err != nil) {
		log.Println(err)
	}
	newPost.UpdatedAt = time.Now()
	
	id, _ := strconv.Atoi(c.Param("id"))
	ok := p.store.Update(uint(id), &newPost)
	if ok {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusNotFound)
	}
}


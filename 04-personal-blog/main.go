package main

import (
	"personal-blog/internal/db"
	"personal-blog/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitDB()
	server.Init(r)
	r.Run()
}
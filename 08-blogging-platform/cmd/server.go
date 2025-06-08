package main

import (
	"blogging-platform/config"
	"blogging-platform/internal/db"
	routes "blogging-platform/internal/route"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	defer db.DB.Close()
	routes.SetupRoutes()
}

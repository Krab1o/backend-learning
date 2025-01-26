package main

import (
	"blogging-platform/config"
	"blogging-platform/internal/db"
	"blogging-platform/internal/routes"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	defer db.DB.Close()
	routes.SetupRoutes()
}

package db

import (
	"blogging-platform/config"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	dbConfig := config.LoadDBConfig()
	connectURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	var err error
	DB, err = sql.Open("pgx", connectURL)
	if (err != nil) {
		log.Fatalf("Init DB: %v", err)
	}
	err = DB.Ping(); if err != nil {
		log.Printf("DB Ping failed: %v", err)
	}
	initStmt, err := os.ReadFile("init.pgsql")
	if (err != nil) {
		log.Fatalf("Init DB: %v", err)
	}
	_, err = DB.Exec(string(initStmt))
	if (err != nil) {
		log.Fatalf("Init DB: %v", err)
	}
}
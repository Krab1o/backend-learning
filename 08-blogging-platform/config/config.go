package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	log.Println(os.Getwd())
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Env loading: %v", err)
	}
}
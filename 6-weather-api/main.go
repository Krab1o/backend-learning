package main

import (
	"context"
	"log"
	"net/http"
	"weather-api/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func main() {
	godotenv.Load()

	Client = redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // No password set
        DB:		  0,  // Use default DB
        Protocol: 2,  // Connection protocol
    })

	if _, err := Client.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
	
	http.HandleFunc("/weather/", handlers.WeatherHandler)
	http.ListenAndServe(":8080", nil)
}
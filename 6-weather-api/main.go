package main

import (
	"net/http"
	"weather-api/internal/handlers"
	"weather-api/internal/redis"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	redis.RedisInit()
	
	http.HandleFunc("/weather/", handlers.WeatherHandler)
	http.ListenAndServe(":8080", nil)
}
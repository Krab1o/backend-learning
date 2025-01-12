package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"weather-api/internal/api"
	"weather-api/internal/data"
	"weather-api/internal/redis"
)

func prettifyJson(w http.ResponseWriter, d *data.Response){
	prettyJSON, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		http.Error(w, "Failed to generate JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyJSON)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city :=	r.URL.Path[len("/weather/"):]
	
	data, ok := redis.GetWithoutContext(city)
	if !ok {
		data = api.WeatherRequest(city)
		redis.SetWithoutContext(city, data, 1 * time.Minute)
	}
	
	prettifyJson(w, data)
}
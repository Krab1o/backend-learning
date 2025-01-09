package handlers

import (
	"encoding/json"
	"net/http"
	"weather-api/internal/data"
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
	
	data := weatherRequest(city)
	prettifyJson(w, data)
}
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"weather-api/internal/data"
)

func weatherRequest(city string) *data.Response {
	apikey := os.Getenv("API_KEY")
	u, _ := url.Parse(data.UrlRequest + city)
	quers := u.Query()
	quers.Set("key", apikey)
	quers.Set("contentType", "json")
	quers.Set("unitGroup", "metric")
	u.RawQuery = quers.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	d := &data.Response{}
	json.NewDecoder(resp.Body).Decode(d)
	return d
}
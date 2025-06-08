package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var url string

func makeRequest(proxied *http.Request) *http.Response {
	
	fmt.Println(url + proxied.RequestURI)
	request, err := http.NewRequest(proxied.Method, "https://" + url + proxied.RequestURI, proxied.Body)
	if (err != nil) {
		log.Println(err)
	}
	response, err := http.DefaultClient.Do(request)
	if (err != nil) {
		log.Println(err)
	}
	return response
}

func proxyHandler(c *gin.Context) {
	response := makeRequest(c.Request)
	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)
	c.JSON(http.StatusOK, gin.H(resp))
}

func Init(port string, origin string) {
	url = origin
	r := gin.Default()
	
	r.Any("/*any", proxyHandler)
	
	r.Run(fmt.Sprintf(":%s", port))
}
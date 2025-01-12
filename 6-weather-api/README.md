# Weather API

Project idea link: https://roadmap.sh/projects/weather-api-wrapper-service

Created mainly on educational purposes.

---

Project that implements querying to [weather API](https://www.visualcrossing.com/weather-api) 
and caches its result with Redis.

It still misses UI interface and now you can only watch the result of the query in browser.

# Setting Up

1. Clone repository
2. Type `go run main.go` or `air` (if you're using [air](https://github.com/air-verse/air))
3. Query `localhost:8080/weather/<city>` so you can get weather data.

## Usage

Basically, you can only access it using web-query (curl or web-browser search bar) using 
example previously mentioned URL.

## Further Plans

- [ ] Add web-site UI (probably using htmx)
- [ ] Dockerize all services
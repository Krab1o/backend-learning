package data

import "time"

const UrlRequest = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"

type Response struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        int     `json:"tzoffset"`
	Days            []struct {
		Datetime       string   `json:"datetime"`
		DatetimeEpoch  int      `json:"datetimeEpoch"`
		Tempmax        float64  `json:"tempmax"`
		Tempmin        int      `json:"tempmin"`
		Temp           int      `json:"temp"`
		Feelslikemax   int      `json:"feelslikemax"`
		Feelslikemin   float64  `json:"feelslikemin"`
		Feelslike      float64  `json:"feelslike"`
		Dew            float64  `json:"dew"`
		Humidity       float64  `json:"humidity"`
		Precip         int      `json:"precip"`
		Precipprob     int      `json:"precipprob"`
		Precipcover    int      `json:"precipcover"`
		Preciptype     any      `json:"preciptype"`
		Snow           int      `json:"snow"`
		Snowdepth      float64  `json:"snowdepth"`
		Windgust       float64  `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        float64  `json:"winddir"`
		Pressure       float64  `json:"pressure"`
		Cloudcover     float64  `json:"cloudcover"`
		Visibility     float64  `json:"visibility"`
		Solarradiation float64  `json:"solarradiation"`
		Solarenergy    float64  `json:"solarenergy"`
		Uvindex        int      `json:"uvindex"`
		Severerisk     int      `json:"severerisk"`
		Sunrise        string   `json:"sunrise"`
		SunriseEpoch   int      `json:"sunriseEpoch"`
		Sunset         string   `json:"sunset"`
		SunsetEpoch    int      `json:"sunsetEpoch"`
		Moonphase      float64  `json:"moonphase"`
		Conditions     string   `json:"conditions"`
		Description    string   `json:"description"`
		Icon           string   `json:"icon"`
		// Stations       []string `json:"stations"`
		Source         string   `json:"source"`
		Hours          []struct {
			Datetime         string    `json:"datetime"`
			DatetimeEpoch    int       `json:"datetimeEpoch"`
			Temp             float64   `json:"temp"`
			Feelslike        float64   `json:"feelslike"`
			Humidity         float64   `json:"humidity"`
			Dew              float64   `json:"dew"`
			Precip           int       `json:"precip"`
			Precipprob       int       `json:"precipprob"`
			Snow             int       `json:"snow"`
			Snowdepth        float64   `json:"snowdepth"`
			Preciptype       any       `json:"preciptype"`
			Windgust         int       `json:"windgust"`
			Windspeed        float64   `json:"windspeed"`
			Winddir          float64   `json:"winddir"`
			Pressure         int       `json:"pressure"`
			Visibility       float64   `json:"visibility"`
			Cloudcover       int       `json:"cloudcover"`
			Solarradiation   int       `json:"solarradiation"`
			Solarenergy      float64   `json:"solarenergy"`
			Uvindex          int       `json:"uvindex"`
			Severerisk       int       `json:"severerisk"`
			Conditions       string    `json:"conditions"`
			Icon             string    `json:"icon"`
			// Stations         []string  `json:"stations"`
			Source           string    `json:"source"`
			DatetimeInstance time.Time `json:"datetimeInstance"`
		} `json:"hours"`
		DatetimeInstance time.Time `json:"datetimeInstance"`
	} `json:"days"`
	Alerts   []any `json:"alerts"`
	// Stations struct {
	// 	F0644 struct {
	// 		Distance     int     `json:"distance"`
	// 		Latitude     float64 `json:"latitude"`
	// 		Longitude    float64 `json:"longitude"`
	// 		UseCount     int     `json:"useCount"`
	// 		ID           string  `json:"id"`
	// 		Name         string  `json:"name"`
	// 		Quality      int     `json:"quality"`
	// 		Contribution int     `json:"contribution"`
	// 	} `json:"F0644"`
	// } `json:"stations"`
	CurrentConditions struct {
		Datetime       string   `json:"datetime"`
		DatetimeEpoch  int      `json:"datetimeEpoch"`
		Temp           float64  `json:"temp"`
		Feelslike      float64  `json:"feelslike"`
		Humidity       float64  `json:"humidity"`
		Dew            float64  `json:"dew"`
		Precip         int      `json:"precip"`
		Precipprob     int      `json:"precipprob"`
		Snow           int      `json:"snow"`
		Snowdepth      int      `json:"snowdepth"`
		Preciptype     any      `json:"preciptype"`
		Windgust       float64  `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        int      `json:"winddir"`
		Pressure       int      `json:"pressure"`
		Visibility     any      `json:"visibility"`
		Cloudcover     float64  `json:"cloudcover"`
		Solarradiation int      `json:"solarradiation"`
		Solarenergy    int      `json:"solarenergy"`
		Uvindex        int      `json:"uvindex"`
		Conditions     string   `json:"conditions"`
		Icon           string   `json:"icon"`
		Stations       []string `json:"stations"`
		Source         string   `json:"source"`
		Sunrise        string   `json:"sunrise"`
		SunriseEpoch   int      `json:"sunriseEpoch"`
		Sunset         string   `json:"sunset"`
		SunsetEpoch    int      `json:"sunsetEpoch"`
		Moonphase      float64  `json:"moonphase"`
	} `json:"currentConditions"`
}
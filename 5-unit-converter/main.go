package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Convertion int

const (
    Length Convertion = iota
    Weight
    Temperature
)

type Data struct {
    ConvertionType  Convertion
    StartValue      float64
    StartUnit       string
    FinalUnit       string
    Result          float64
}

func (c Convertion) String() string {
    switch c {
    case 0:
        return "length"
    case 1:
        return "weight"
    case 2:
        return "temperature"
    default:
        return "unknown"
    }
}

var lengths = map[string]float64{
    "millimeter": 0.001,
    "centimeter" : 0.01,
    "meter": 1,
    "kilometer": 1000,
    "inch" : 0.0254,
    "foot" : 0.3048,
    "yard" : 0.9144,
    "mile" : 1609.34,
}

var weights = map[string]float64{
    "milligram" : 0.001,
    "gram" : 1,
    "kilogram": 1000,
    "ounce": 28.3495,
    "pound": 453.592,
}

type TempUnit int

const (
    Celsius TempUnit = iota
    Fahrenheit
    Kelvin
)

func (c TempUnit) String() string {
    switch c {
    case 0:
        return "celsius"
    case 1:
        return "fahrenheit"
    case 2:
        return "kelvin"
    default:
        return "unknown"
    }
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    switch status {
    case http.StatusNotFound:
        notFound, _ := os.ReadFile("static/html/not_found.html")
        fmt.Fprint(w, string(notFound))
    }
}

func converterHandler(w http.ResponseWriter, r *http.Request) {
    convertType := r.URL.Path[len("/convert/"):]
    file, err := os.ReadFile("static/html/convert_" + convertType + ".html")
    if (err != nil) {
        errorHandler(w, r, http.StatusNotFound)
    }
    fmt.Fprint(w, string(file))    
}

func parseData(bytes []byte) *Data {
    str := strings.Split(string(bytes), "&")

    value, err := strconv.ParseFloat(str[0][len("value="):], 64)
    if (err != nil) {
        log.Println(err)
    }

    startUnit := str[1][len("converted="):]
    finalUnit := str[2][len("converting="):]

    return &Data {
        StartValue: value,
        StartUnit:  startUnit,
        FinalUnit:  finalUnit,
    }
}


func renderTemplate(w http.ResponseWriter, data *Data) {
    t, err := template.ParseFiles("template/result.html")
    if (err != nil) {
        log.Println(err)
    }
    err = t.Execute(w, data)
    if (err != nil) {
        log.Println(err)
    }
}

func calculateTemperature(data *Data) {
    switch data.StartUnit {
    case data.FinalUnit:
        data.Result = data.StartValue
    case Celsius.String():
        log.Println("===celsius===")
        switch data.FinalUnit {
        // celsius to fahrenheit
        case Fahrenheit.String():
            data.Result = (data.StartValue * 9 / 5) + 32 
        // celsius to kelvin
        case Kelvin.String():
            data.Result = data.StartValue + 273.15
        }
    case Fahrenheit.String():
        switch data.FinalUnit {
        // fahrenheit to celsius
        case Celsius.String():
            data.Result = (data.StartValue - 32) * 5 / 9
        // fahrenheit to kelvin
        case Kelvin.String():
            data.Result = (data.StartValue - 32) * 5 / 9 + 273.15
        }
    case Kelvin.String():
        switch data.FinalUnit {
        //kelvin to celsius
        case Celsius.String():
            data.Result = data.StartValue - 273.15
        //kelvin to fahrenheit
        case Fahrenheit.String():
            data.Result = (data.StartValue - 273.15) * 9 / 5 + 32 
        }
    default:
        data.Result = 100
    }
    log.Println(data.Result)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
    convertType := r.URL.Path[len("/calculate/"):]
    bytes, err := io.ReadAll(r.Body)
    if err != nil {
        log.Println(err)
    }

    data := parseData(bytes)
    switch convertType {
    case "length":
        data.ConvertionType = Length
        data.Result = data.StartValue * lengths[data.StartUnit] / lengths[data.FinalUnit]
    case "weight":
        data.ConvertionType = Weight
        data.Result = data.StartValue * weights[data.StartUnit] / weights[data.FinalUnit]
    case "temperature":
        data.ConvertionType = Temperature
        calculateTemperature(data)
    }
    log.Println(data)
    renderTemplate(w, data)
}

func main() {
    
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/calculate/", calculateHandler)
    http.HandleFunc("/convert/", converterHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}



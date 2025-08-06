package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"
)

type WeatherData struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Name    string    `json:"name"`
}

type Weather struct {
	Description string `json:"description"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

func main() {
	for {
		var weatherID string = "745042"

		baseUrl := "https://api.openweathermap.org/data/2.5/weather?id=" + weatherID + "&appid=5796abbde9106b7da4febfae8c44c232"
		resp, err := http.Get(baseUrl)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var weatherData WeatherData
		json.Unmarshal(body, &weatherData)

		temp := weatherData.Main.Temp - 273.15
		wind_speed := weatherData.Wind.Speed
		fmt.Printf(
			"\rCountry: %s | Weather: %s | Temperature: %.2f°C | Wind Speed: %.2f m/s",
			weatherData.Name,
			strings.Title(strings.ToLower(weatherData.Weather[0].Description)),
			math.Round(temp),
			wind_speed,
		)

		time.Sleep(2 * time.Second)
	}

}

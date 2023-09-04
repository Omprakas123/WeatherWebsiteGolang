package main

import (
	// "bufio"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	// "os"
)

// Assuming you have JSON data from a weather API like this: so we have to convert that data in structure format so here we made that json data format
type WeatherData struct {
	Coord      coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        uint16    `json:"cod"`
}

type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	Temp_min   float64 `json:"temp_min"`
	Temp_max   float64 `json:"temp_max"`
	Pressure   int64   `json:"pressure"`
	Humidity   int64   `json:"humidity"`
	Sea_level  int64   `json:"sea_level"`
	Grnd_level int64   `json:"grnd_level"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
type Clouds struct {
	All int64 `json:"all"`
}

type Sys struct {
	Type    int32  `json:"type"`
	Id      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you City Name  and State Code So that I can display all weather data of that city:->")
	CityName, _ := reader.ReadString('\n')
	StateCode, _ := reader.ReadString('\n')
	url := "https://api.openweathermap.org/data/2.5/weather?IN&units=metric&appid=88579dae9dda5c199f158e88cdc480e9&q="
	// Concatenate the query parameter to the URL
	fullURL := url + CityName[:len(CityName)-1] + "," + StateCode[:len(StateCode)-1] + "+91"
	// Make an HTTP GET request to a URL
	fmt.Println(StateCode)
	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	AllWeatherData := string(body)
	jsonData := []byte(AllWeatherData)
	// Create a variable of the WeatherData struct type
	var weatherData WeatherData

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(jsonData, &weatherData); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Here you can get weather information of that particular city that you have search in this website")
	fmt.Println("CityName->", weatherData.Name)
	fmt.Println("CityWeatherType->", weatherData.Weather[0].Main)
	fmt.Println("CityLongitude->", weatherData.Coord.Lon)
	fmt.Println("CityLatitude->", weatherData.Coord.Lat)
	fmt.Println("CityMinimumTemp->", weatherData.Main.Temp_min, ".C")
	fmt.Println("CityMaximumTemp->", weatherData.Main.Temp_max, ".C")
	fmt.Println("Presure->", weatherData.Main.Pressure)
	fmt.Println("Humidity->", weatherData.Main.Humidity)

}

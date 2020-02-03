package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

var apiKey = "xxxx"

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Input file is missing.")
		os.Exit(1)
	}

	city := args[0]

	w, err := owm.NewForecast("5", "C", "EN", apiKey) // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Fatalln(err)
	}

	_ = w.DailyByName(city, 1)

	data := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	if len(data.List) < 1 {
		fmt.Println("Result not found, probably there is no such city. Check for typos.")
		os.Exit(1)
	}
	
	if data.List[0].Main.TempMin < 15 {
		fmt.Println("hoodie")
	} else {
		fmt.Println("no hoodie")
	}
}

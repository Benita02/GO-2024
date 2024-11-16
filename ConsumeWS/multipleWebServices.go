package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var apis = map[int]string{
	1: "http://data.fixer.io/api/latest?access_key=34f101b1e43c596ecdc669eeb73e27e0",
	2: "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=d317d04187dd928b884fd1b6697c05b",
}

func fetchData(apiID int) {
	url, ok := apis[apiID]
	if !ok {
		log.Fatalf("Invalid API ID: %d", apiID)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data from API %d: %v", apiID, err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding response from API %d: %v", apiID, err)
	}

	switch apiID {
	case 1: // Fixer.io API
		handleFixerResponse(result)
	case 2: // OpenWeatherMap API
		handleWeatherResponse(result)
	default:
		log.Fatalf("Unknown API ID: %d", apiID)
	}
}

func handleFixerResponse(result map[string]interface{}) {
	if success, ok := result["success"].(bool); ok && success {
		if rates, ok := result["rates"].(map[string]interface{}); ok {
			if usd, ok := rates["USD"].(float64); ok {
				fmt.Printf("USD Exchange Rate: %.2f\n", usd)
				return
			}
		}
		log.Println("Unexpected rates format in Fixer.io response.")
	} else if errorInfo, ok := result["error"].(map[string]interface{}); ok {
		fmt.Println("Fixer.io Error:", errorInfo["info"])
	} else {
		log.Println("Unexpected Fixer.io error format.")
	}
}

func handleWeatherResponse(result map[string]interface{}) {
	if main, ok := result["main"].(map[string]interface{}); ok {
		if temp, ok := main["temp"].(float64); ok {
			fmt.Printf("Temperature in Singapore: %.2f°C\n", temp-273.15) // Convert Kelvin to Celsius
			return
		}
		log.Println("Unexpected temperature format in OpenWeatherMap response.")
	} else if message, ok := result["message"].(string); ok {
		fmt.Println("OpenWeatherMap Error:", message)
	} else {
		log.Println("Unexpected OpenWeatherMap error format.")
	}
}

func main() {
	fetchData(1)
	fetchData(2)
}

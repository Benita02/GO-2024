package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var apis = map[int]string{
	1: "http://data.fixer.io/api/latest?access_key=34f101b1e43c596ecdc669eeb73e27e0",
	2: "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=6b4d7ef379aea53ad8f4e224e5be2161",
}

// Result struct to encapsulate API responses and potential errors
type Result struct {
	APIID int
	Data  map[string]interface{}
	Error error
}

func fetchData(apiID int, ch chan<- Result) {
	url, ok := apis[apiID]
	if !ok {
		ch <- Result{APIID: apiID, Error: fmt.Errorf("invalid API ID: %d", apiID)}
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{APIID: apiID, Error: fmt.Errorf("error fetching data from API %d: %v", apiID, err)}
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		ch <- Result{APIID: apiID, Error: fmt.Errorf("error decoding response from API %d: %v", apiID, err)}
		return
	}

	ch <- Result{APIID: apiID, Data: result}
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
	// Create a channel to collect results from Goroutines
	results := make(chan Result, len(apis))

	// Launch Goroutines for each API
	for apiID := range apis {
		go fetchData(apiID, results)
	}

	// Collect and process results
	for i := 0; i < len(apis); i++ {
		result := <-results
		if result.Error != nil {
			log.Printf("Error from API %d: %v\n", result.APIID, result.Error)
			continue
		}

		switch result.APIID {
		case 1:
			handleFixerResponse(result.Data)
		case 2:
			handleWeatherResponse(result.Data)
		}
	}
}

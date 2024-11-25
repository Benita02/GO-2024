package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var apis = map[int]string{
	1: "http://data.fixer.io/api/latest?access_key=<access_key>",
	2: "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=<app_id>",
}

type Result struct {
	APIID int
	Data  interface{}
	Error error
}

func fetchData(apiID int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	url, ok := apis[apiID]
	if !ok {
		results <- Result{APIID: apiID, Error: fmt.Errorf("invalid API ID: %d", apiID)}
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		results <- Result{APIID: apiID, Error: fmt.Errorf("failed to fetch API %d: %v", apiID, err)}
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		results <- Result{APIID: apiID, Error: fmt.Errorf("failed to decode API %d: %v", apiID, err)}
		return
	}

	results <- Result{APIID: apiID, Data: data}
}

func processResult(result Result) {
	if result.Error != nil {
		log.Printf("Error from API %d: %v", result.APIID, result.Error)
		return
	}

	switch result.APIID {
	case 1:
		handleFixerResponse(result.Data)
	case 2:
		handleWeatherResponse(result.Data)
		// default:
		// 	log.Printf("Unknown API ID: %d", result.APIID)
		// } THIS IS REDUNDANT TO ME, ALREADY CHECKED THIS AT THE VERY BEGINNING
	}
}

func handleFixerResponse(data interface{}) {
	result, ok := data.(map[string]interface{})
	if !ok {
		log.Println("Invalid Fixer.io data format.")
		return
	}

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

func handleWeatherResponse(data interface{}) {
	result, ok := data.(map[string]interface{})
	if !ok {
		log.Println("Invalid OpenWeatherMap data format.")
		return
	}

	if main, ok := result["main"].(map[string]interface{}); ok {
		if temp, ok := main["temp"].(float64); ok {
			fmt.Printf("Temperature in Singapore: %.2f°C\n", temp-273.15)
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
	results := make(chan Result, len(apis))
	var wg sync.WaitGroup

	for apiID := range apis {
		wg.Add(1)
		go fetchData(apiID, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		processResult(result)
	}
}

// God abeg my head dey pain me. I nor do again
// READING ABOUT RESTAPIs

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	apis = map[int]string{
		1: "http://data.fixer.io/api/latest?access_key=<access_key>",
		2: "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=<app_id>",
	}
)

type APIResult struct {
	ID    int
	Value interface{}
	Error error
}

func fetchData(apiID int, ch chan<- APIResult, wg *sync.WaitGroup) {
	defer wg.Done()

	url, ok := apis[apiID]
	if !ok {
		ch <- APIResult{ID: apiID, Error: fmt.Errorf("invalid API ID: %d", apiID)}
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- APIResult{ID: apiID, Error: fmt.Errorf("failed to fetch data: %w", err)}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- APIResult{ID: apiID, Error: fmt.Errorf("failed to read response: %w", err)}
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		ch <- APIResult{ID: apiID, Error: fmt.Errorf("failed to parse JSON: %w", err)}
		return
	}

	var apiResult interface{}
	switch apiID {
	case 1: // Currency API
		if success, ok := result["success"].(bool); ok && success {
			if rates, ok := result["rates"].(map[string]interface{}); ok {
				apiResult = rates["USD"]
			}
		} else {
			if errorInfo, ok := result["error"].(map[string]interface{}); ok {
				apiResult = errorInfo["info"]
			}
		}
	case 2: // Weather API
		if main, ok := result["main"].(map[string]interface{}); ok {
			apiResult = main["temp"]
		} else {
			apiResult = result["message"]
		}
	default:
		ch <- APIResult{ID: apiID, Error: fmt.Errorf("unsupported API ID: %d", apiID)}
		return
	}

	ch <- APIResult{ID: apiID, Value: apiResult}
}

func main() {
	results := make(chan APIResult)
	var wg sync.WaitGroup

	// Launch fetch tasks for each API
	for apiID := range apis {
		wg.Add(1)
		go fetchData(apiID, results, &wg)
	}

	// Close the results channel once all tasks complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print results
	for result := range results {
		if result.Error != nil {
			log.Printf("API %d error: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("API %d result: %v\n", result.ID, result.Value)
		}
	}
}

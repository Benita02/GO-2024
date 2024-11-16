package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Success   bool               `json:"success"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

type APIError struct {
	Success bool `json:"success"`
	Error   struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error"`
}

func main() {
	url := "http://data.fixer.io/api/latest?access_key=34f101b1e43c596ecdc669eeb73e27e0"

	result, err := fetchExchangeRates(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print exchange rates
	for currency, rate := range result.Rates {
		fmt.Printf("%s: %f\n", currency, rate)
	}

	fmt.Println("Done")
}

// fetchExchangeRates fetches the exchange rates and handles success or failure responses.
func fetchExchangeRates(url string) (*Result, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// Decode JSON response directly
	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	// Check if the API response indicates failure
	if !result.Success {
		var apiError APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return nil, fmt.Errorf("failed to decode API error: %w", err)
		}
		return nil, fmt.Errorf("API error: %s", apiError.Error.Info)
	}

	return &result, nil
}

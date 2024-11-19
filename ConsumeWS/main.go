package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

type Result struct {
	Success   bool               `json:"success"`
	Rates     map[string]float64 `json:"rates"`
	ErrorInfo APIError           `json:"error"`
}

type APIError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}

func main() {
	url := "http://data.fixer.io/api/latest?access_key=34f101b1e43c596ecdc669eeb73e27e0"
	// read more of this topic on phone

	result, err := fetchExchangeRates(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	printRates(result.Rates)
	fmt.Println("Done")
}

func fetchExchangeRates(url string) (*Result, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	if !result.Success {
		return nil, fmt.Errorf("API error: %s", result.ErrorInfo.Info)
	}

	return &result, nil
}

func printRates(rates map[string]float64) {
	keys := make([]string, 0, len(rates))
	for currency := range rates {
		keys = append(keys, currency)
	}
	sort.Strings(keys)

	for _, currency := range keys {
		fmt.Printf("%s: %f\n", currency, rates[currency])
	}
}

// continued work from phone, nothing to commit
// laptop, 18%, continuing work from phone, nothing to commit

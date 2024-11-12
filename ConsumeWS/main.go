package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://data.fixer.io/api/latest?access_key=<access_key>"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(body))
	fmt.Println("Done")
}

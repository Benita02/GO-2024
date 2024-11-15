package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://data.fixer.io/api/latest?access_key=34f101b1e43c596ecdc669eeb73e27e0"

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

//learning optimized version of how to access the fixer api on phone since no electricity to charge laptop

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	defer fmt.Println("...::: this will be printed at the end of the current function :::...")

	client := http.Client{}
	resp, err := client.Get("https://google.com")
	if err != nil {
		log.Fatalf("client get: %s\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read response body: %s\n", err)
	}
	defer resp.Body.Close()

	// multiple defers will be executed in FILO order

	log.Println(string(body))
}

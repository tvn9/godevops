package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// http get request example

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./httpget <url>")
		os.Exit(1)
	}

	input := args[1]

	// validate for valid url
	URL, err := url.ParseRequestURI(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("URL.String()", URL.String())
	res, err := http.Get(input)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP status code: %d\nBody: %s\n", res.StatusCode, body)
}

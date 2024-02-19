package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// http get request example

// {"page":"words","input":"word1","words":["word1"]}

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

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

	if res.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP code %d): %s\n", res.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))
}

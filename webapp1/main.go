package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// json data {"page":"word1","words":["word1"]} to be parsed

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args[1:]
	webUrl := args[0]

	if len(args) < 1 {
		fmt.Println("Usage: ./http-get <url>")
		os.Exit(1)
	}
	_, err := url.ParseRequestURI(webUrl)
	if err != nil {
		log.Fatal(err)
	}

	rep, err := http.Get(webUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer rep.Body.Close()

	body, err := io.ReadAll(rep.Body)
	if err != nil {
		log.Fatal(err)
	}

	if rep.StatusCode != 200 {
		fmt.Printf("invalid output (HTTP code %d)\n", rep.StatusCode)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON Parsed")
	fmt.Printf("Page: %s\n", words.Page)
	fmt.Printf("Words: %s\n", words.Words)
}

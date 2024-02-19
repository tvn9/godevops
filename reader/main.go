package main

import (
	"fmt"
	"io"
	"log"
)

// Rewrite the http get client with our own read function

type MySlowReader struct {
	contents string
	pos      int
}

func (r *MySlowReader) Read(b []byte) (n int, err error) {
	if r.pos+1 <= len(r.contents) {
		n := copy(b, r.contents[r.pos:r.pos+1])
		r.pos++
		return n, nil
	}
	return n, io.EOF
}

func main() {

	mySlowReader := &MySlowReader{
		contents: "Hello World!",
	}

	out, err := io.ReadAll(mySlowReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output %s", out)
}

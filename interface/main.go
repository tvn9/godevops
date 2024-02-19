package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	Contents string
	Count    int
}

func (m *MySlowReader) Read(b []byte) (n int, err error) {
	if m.Count < len(m.Contents) {
		n := copy(b, m.Contents[m.Count:m.Count+1])
		m.Count++
		return n, nil
	}
	return 0, io.EOF
}

func main() {
	msr := &MySlowReader{
		Contents: "Hello World!",
	}

	out, err := io.ReadAll(msr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %s - Byte Count: %d\n", string(out), msr.Count)
}

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "Hello, Go!"

	reader := strings.NewReader(s)
	buf := make([]byte, 4)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}
}

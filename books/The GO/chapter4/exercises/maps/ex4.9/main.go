package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordfreq("input.txt")
}

func wordfreq(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("File %s does not exist", err)
		return
	}

	scanner := bufio.NewScanner(file)
	counts := map[string]int{}
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " ")
		for _, text := range lines {
			counts[strings.ToLower(text)]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

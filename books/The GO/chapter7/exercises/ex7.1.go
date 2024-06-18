package exercises

import (
	"bufio"
	"os"
)

type Reader struct {
	Readers *os.File
}

func (r *Reader) WordsCounter() int {
	count := 0
	scanner := bufio.NewScanner(r.Readers)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}

func (r *Reader) LinesCounter() int {
	count := 0
	scanner := bufio.NewScanner(r.Readers)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count++
	}
	return count
}

func ReadFile(filePath string) (*Reader, error) {
	file, err := os.Open("example.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return &Reader{Readers: file}, nil
}

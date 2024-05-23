package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var algorithm string
	flag.StringVar(&algorithm, "algorithm", "", "Specify your name")
	flag.Parse()
	c1 := hashAlgorithm(algorithm, "x")
	c2 := hashAlgorithm(algorithm, "X")
	fmt.Printf("%x\n%x\n", c1, c2)
}

func hashAlgorithm(al string, input string) []byte {
	hashRes := []byte{}
	switch strings.ToLower(al) {
	case "sha384":
		hashArray := sha512.Sum384([]byte(input))
		hashRes = hashArray[:]
	case "sha512":
		hashArray := sha512.Sum512([]byte(input))
		hashRes = hashArray[:]
	default:
		hashArray := sha256.Sum256([]byte(input))
		hashRes = hashArray[:]
	}
	return hashRes
}

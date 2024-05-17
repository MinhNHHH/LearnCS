package main

import (
	"crypto/sha256"
	"fmt"
)

// PopCount returns the population count (number of set bits) of x.
func differBitCount(x, y uint8) int {
	diff := 0
	for i := 0; i < 8; i++ {
		bit1 := (x >> i) & 1
		bit2 := (y >> i) & 1
		if bit1 != bit2 {
			diff++
		}
	}
	return diff
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	differBit := 0
	for i := 0; i < len(c1); i++ {
		differBit += differBitCount(c1[i], c2[i])
	}
	fmt.Println(differBit)
}

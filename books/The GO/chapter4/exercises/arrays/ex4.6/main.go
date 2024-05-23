package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(squashSpace([]byte("Hello,     世界!  This    is a    test.")))
}

func squashSpace(s []byte) []byte {
	n := 0
	spaceFound := false
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			if !spaceFound {
				s[n] = ' '
				n++
				spaceFound = true
			}
		} else {
			s[n] = s[i]
			n++
			spaceFound = false
		}
	}
	return s[:n] // Reslice to the new length
}

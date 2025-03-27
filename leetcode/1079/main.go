package main

import "fmt"

func backtrack(count map[rune]int) int {
	sum := 0
	for ch, c := range count {
		if c > 0 {
			// Chose the letter
			count[ch]--
			sum++
			sum += backtrack(count)
			// backtrack
			count[ch]++
		}
	}
	return sum
}

func numTilePossibilities(tiles string) int {
	count := map[rune]int{}
	for _, t := range tiles {
		count[t]++
	}

	return backtrack(count)
}

func main() {
	fmt.Println(numTilePossibilities("AAB"))
}

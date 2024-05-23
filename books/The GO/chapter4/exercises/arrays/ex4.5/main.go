package main

import "fmt"

func main() {
	originalSlice := []string{"apple", "banana", "apple", "orange", "banana", "grape"}

	fmt.Println(removeDuplicate(originalSlice))
}

func removeDuplicate(listString []string) []string {
	setString := map[string]bool{}
	res := []string{}
	for _, str := range listString {
		if !setString[str] {
			setString[str] = true
			res = append(res, str)
		}
	}

	return res
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var ages map[string]int
	if age, ok := ages["bob"]; !ok {
		ages["bob"] = age
	}
	// The order of map iteration is unspecified.
	// In practice, the order is random, varying from one execution to the next.
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func enumerateKeyValueRandomOrder(dict map[string]int) {
	// The order of map iteration is unspecified.
	// In practice, the order is random, varying from one execution to the next.
	for name, age := range dict {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func enumerateKeyvalueOrder(dict map[string]int) {

	var names []string
	for name := range dict {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, dict[name])
	}
}

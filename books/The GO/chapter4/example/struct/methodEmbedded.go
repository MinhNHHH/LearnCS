package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "says hello")
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	d := Dog{
		Animal: Animal{
			Name: "Buddy",
		},
		Breed: "Golden Retriever",
	}
	d.Speak() // Output: Buddy says hello
}

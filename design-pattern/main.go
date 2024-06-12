package main

import "fmt"

// Product interface
type Shape interface {
	Draw()
}

// Concrete product classes
type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Circle.draw()")
}

type Square struct{}

func (s Square) Draw() {
	fmt.Println("Square.draw()")
}

// Factory interface
type ShapeFactory interface {
	CreateShape() Shape
}

// Concrete factory classes
type CircleFactory struct{}

func (cf CircleFactory) CreateShape() Shape {
	return Circle{}
}

type SquareFactory struct{}

func (sf SquareFactory) CreateShape() Shape {
	return Square{}
}

// Client code
func DrawShape(factory ShapeFactory) {
	shape := factory.CreateShape()
	shape.Draw()
}

// Example usage
func main() {
	DrawShape(CircleFactory{}) // Output: Circle.draw()
	DrawShape(SquareFactory{}) // Output: Square.draw()
}

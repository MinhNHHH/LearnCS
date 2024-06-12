package main

import "fmt"

type ColorImplementor interface {
	filColor() string
}

type ShapeImplementor interface {
	draw() string
	setColor(ColorImplementor)
}

type ColorShaped struct {
	shape string
	color ColorImplementor
}

type Color struct {
	color string
}

func (c *Color) filColor() string {
	return c.color
}

func (cs *ColorShaped) draw() string {
	return fmt.Sprintf("Draw %s with color %s", cs.shape, cs.color.filColor())
}

func (cs *ColorShaped) setColor(c ColorImplementor) {
	cs.color = c
}

func main() {
	// Create a circle shape
	circle := ColorShaped{shape: "circle"}

	// Set the color of the circle to red
	circle.setColor(&Color{color: "red"})

	// Draw the circle
	fmt.Println(circle.draw())
}

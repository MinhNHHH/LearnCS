package main

import (
	"fmt"
	"math"
)

type Shape interface {
	GetWidth() float64
	GetRadius() float64
}

type Rectangle struct {
	width float64
}

type Circle struct {
	radius float64
}

func NewRectangle(width float64) *Rectangle {
	return &Rectangle{width: width}
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (r *Rectangle) GetWidth() float64 {
	return r.width
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

type ShapeRectangleAdapter struct {
	sr *Rectangle
}

func (rec *ShapeRectangleAdapter) GetRadius() float64 {
	return rec.sr.GetWidth() * math.Sqrt(2) / 2
}

type ShapeCircleAdapter struct {
	cr *Circle
}

func (rec *ShapeCircleAdapter) GetWidth() float64 {
	return rec.cr.GetRadius() * 2
}

func main() {
	r := NewRectangle(2.0)
	c := NewCircle(2.0)

	ra := &ShapeRectangleAdapter{sr: r}
	ca := &ShapeCircleAdapter{cr: c}

	fmt.Println("Rectangle's diagonal (radius):", ra.GetRadius())
	fmt.Println("Circle's diameter (width):", ca.GetWidth())
}

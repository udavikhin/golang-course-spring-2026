package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printShapesAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}

func main() {
	printShapesAreas([]Shape{
		Circle{
			Radius: 10,
		},
		Rectangle{
			Width:  20,
			Height: 30,
		},
	})
}

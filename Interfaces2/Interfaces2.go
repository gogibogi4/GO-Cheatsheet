package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

//define area() for Rect
func (r Rect) area() float64 {
	return r.Width * r.Height
}

//define area() for Circle
func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//define perim() for Rect
func (r Rect) perim() float64 {
	return 2*r.Width + 2 + r.Height
}

//define perim() for Circle
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.Radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := Rect{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	measure(r)
	measure(c)
}

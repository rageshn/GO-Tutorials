package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	s := square{
		length: 10.00,
		width:  10,
	}
	c := circle{
		radius: 10,
	}

	sa := info(s)
	fmt.Println("Area of square:", sa)
	sc := info(c)
	fmt.Println("Area of circle:", sc)

}

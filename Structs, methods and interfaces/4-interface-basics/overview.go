package main

import (
	"fmt"
	"math"
)

type Square struct {
	sideLength float64
}

func (s Square) GetArea() float64 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float64
}

func (c Circle) GetArea() float64 {
	return c.radius * c.radius * math.Pi
}

func overview() {
	s := Square{
		sideLength: 10,
	}

	c := Circle{
		radius: 10,
	}

	printAreaOfSquare(s)
	printAreaOfCircle(c)

	printAreaOfShape()
}

func printAreaOfSquare(s Square) {
	fmt.Printf("The area of this Square is: %f\n", s.GetArea())
}

func printAreaOfCircle(c Circle) {
	fmt.Printf("The area of this Circle is: %f\n", c.GetArea())
}

// Fill this function with logic that will take in one parameter either a square or a cirlce
// and will print out a message like: "The area of this Square is 100"
func printAreaOfShape() {

}

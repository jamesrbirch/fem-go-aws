package main

import (
	"fmt"
)

const (
	PI = 3.14
)

type Circle struct {
	Radius float64
}

func (circle Circle) Circumference() float64 {
	circumference := (2 * PI * circle.Radius)
	return circumference
}

func (circle Circle) Area() float64 {
	area := PI * (circle.Radius * circle.Radius)
	return area
}

func main() {
	firstCircle := Circle{Radius: 1.5}

	fmt.Println(firstCircle.Circumference())
	fmt.Println(firstCircle.Area())
}

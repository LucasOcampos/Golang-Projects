package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	width  float64
}
type square struct {
	length float64
	width  float64
}

func main() {
	squareShape := square{10, 20}
	triangleShape := triangle{10, 20}

	printArea(squareShape)
	printArea(triangleShape)
}

func (t triangle) getArea() float64 {
	return (t.width * t.height) / 2
}

func (s square) getArea() float64 {
	return s.length * s.width
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

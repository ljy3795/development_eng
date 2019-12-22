package main

import (
	"fmt"
	"math"
)

// Shape interface has a method area(), perimeter()
type Shape interface {
	area() float64
	// perimeter() float64
}

// struct
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// method
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// getArea(s Shape)
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("1")

	c := Circle{x: 0, y: 0, radius: 5}
	r := Rectangle{width: 10, height: 5}

	fmt.Println("Area for Circls is :")
	fmt.Println(c.area())
	// passing an interface(s Shape) but not a specific implementaion (neither Circle nor Rectangle)
	fmt.Println(getArea(c))

	fmt.Println("Area for Rectangle is :")
	fmt.Println(r.area())
	// passing an interface(s Shape) but not a specific implementaion (neither Circle nor Rectangle)
	fmt.Println(getArea(r))

	// EMPTY interface
	var value interface{}
	show(value)

	value = 49
	show(value)

	value = "empty interface"
	show(value)
}

// interface{} type is the interface that has no methods
//  -> paramter with interface{}, we can supply that function with any value
func show(value interface{}) {
	fmt.Printf("(%v, %T)\n", value, value)
}

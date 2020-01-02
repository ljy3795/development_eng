package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
}

func (p *Point) size() float64 {
	return math.Sqrt(p.x*p.x + p.y + p.y)
}

func main() {
	p := &Point{3, 4} // set pointer
	fmt.Println(p)
	fmt.Println(*p)

	p.scale(2)
	fmt.Println(p)
	fmt.Println(p.size())
}

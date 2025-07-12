package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// The Distance method calculates the distance from the current point to another
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)

	fmt.Println(p1.Distance(p2))
}

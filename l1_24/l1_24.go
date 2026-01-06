package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {

	return Point{x: x, y: y}
}

func (p Point) Distance(secondP Point) float64 {
	dist := math.Sqrt(math.Pow(secondP.x-p.x, 2) + math.Pow(secondP.y-p.y, 2))

	return dist
}

func main() {
	p1 := NewPoint(2, 3)
	p2 := NewPoint(4, 7)

	fmt.Printf("Дистанция: %.2f\n", p1.Distance(p2))
}

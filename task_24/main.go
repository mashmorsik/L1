package main

import (
	"fmt"
	"math"
)

/* Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором. */

func main() {
	a, b := NewPoint(2, 4), NewPoint(4, 2)

	fmt.Println(a.GetDistance(b))
	fmt.Println(b.GetDistance(a))
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetDistance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2))
}

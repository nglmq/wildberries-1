package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point { // конструктор
	return &Point{
		x: x,
		y: y,
	}
}

func checkDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2)) // расстояние между точками
}

func main() {
	a := NewPoint(-1, 3)
	b := NewPoint(6, 2)

	distance := checkDistance(a, b)
	fmt.Println(distance)
}

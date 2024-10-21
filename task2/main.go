package main

import (
	"fmt"
	"math"
	"sort"
)

const pi float64 = 3.14

type Shaper interface {
	Area() float64
	Perimeter() float64
	IsLargerThan(Shaper) bool
}

// Circle /////////////////////////////////////
type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return pi * math.Pow(c.radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * pi * c.radius
}

func (c *Circle) IsLargerThan(other Shaper) bool {
	if c.Area() > other.Area() {
		return true
	}

	return false
}

// Triangle ///////////////////////////////////////////////////
type Triangle struct {
	a, b, c float64
}

func (tr *Triangle) Area() float64 {
	p := (tr.a + tr.b + tr.c) / 2
	return math.Sqrt(p * (p - tr.a) * (p - tr.b) * (p - tr.c))
}

func (tr *Triangle) Perimeter() float64 {
	return tr.a + tr.b + tr.c
}

func (tr *Triangle) IsLargerThan(other Shaper) bool {
	if tr.Area() > other.Area() {
		return true
	}

	return false
}

// Rectangle //////////////////////////////////////////////////////////////
type Rectangle struct {
	a, b float64
}

func (rc Rectangle) Area() float64 { // ресивер убрал передачу по указателю для вывода результатов функций
	return rc.a * rc.b
}

func (rc Rectangle) Perimeter() float64 {
	return 2 * (rc.a + rc.b)
}

func (rc Rectangle) IsLargerThan(other Shaper) bool {
	if rc.Area() > other.Area() {
		return true
	}

	return false
}

func (rc Rectangle) TransformToSquare() float64 {
	minVal := min(rc.a, rc.b)
	return math.Pow(minVal, 2)
}

// SortShapes Function//////////////////////////////////////////////////////////////
func SortShapes(shapes []Shaper) {
	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].Area() < shapes[j].Area()
	})
}

// FilterShapes Function/////////////////////////////////////////////////////////////
func FilterShapes(shapes []Shaper, minArea float64) []Shaper {
	SortShapes(shapes)
	for i := range shapes {
		if shapes[i].Area() > minArea {
			return shapes[i:]
		}
	}

	return []Shaper{}
}

// min function////////////////////////////////////////////
func min(a, b float64) float64 {
	if a == b {
		return a
	} else if a < b {
		return a
	}

	return b

}

func main() {
	rec := Rectangle{6, 4}
	triang := Triangle{3, 4, 5}
	fmt.Println(rec.Area(), rec.Perimeter(), rec.IsLargerThan(&Rectangle{5, 5}), rec.TransformToSquare()) //??? Ресивер передаю указатель в методе для сравнения фигур, почему аргумент other просит по указателю
	fmt.Println(FilterShapes([]Shaper{Rectangle{2, 3}, Rectangle{2, 2}, Rectangle{5, 4}, Rectangle{8, 8}}, 15))
	fmt.Println(triang.Area())
}

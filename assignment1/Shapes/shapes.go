package Shapes

import "math"

type Shape interface {
	CalculateArea() float64
	CalculatePerimetr() float64
}

type Rectangle struct {
	A float64
	B float64
}
type Circle struct {
	R float64
}
type Square struct {
	A float64
}
type Triangle struct {
	A float64
	B float64
	C float64
}

// Areas
func (rect Rectangle) CalculateArea() (area float64) {
	area = rect.A * rect.B
	return
}
func (circle Circle) CalculateArea() (area float64) {
	area = math.Pi * circle.R * circle.R
	return
}
func (square Square) CalculateArea() (area float64) {
	area = square.A * square.A
	return
}
func (tri Triangle) CalculateArea() (area float64) {
	semi := (tri.A + tri.B + tri.C) / 2
	area = math.Sqrt(semi * (semi - tri.A) * (semi - tri.B) * (semi - tri.C))
	return
}

// Prerimetrs
func (rect Rectangle) CalculatePerimetr() (perimetr float64) {
	perimetr = 2 * (rect.A + rect.B)
	return
}
func (circle Circle) CalculatePerimetr() (perimetr float64) {
	perimetr = 2 * math.Pi * circle.R
	return
}
func (square Square) CalculatePerimetr() (perimetr float64) {
	perimetr = 4 * square.A
	return
}
func (tri Triangle) CalculatePerimetr() (perimetr float64) {
	perimetr = tri.A + tri.B + tri.C
	return
}

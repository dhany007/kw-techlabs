package interface4

import (
	"fmt"
	"math"
)

// interface = tipe data untuk sekumpulan method atau tipe data

// shape disini mempunyai 2 method
type Shape interface {
	Area() float64
	Parimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// membuat nilai rectangle
func NewRectangle(width, heigt float64) Shape {
	return Rectangle{
		Width:  width,
		Height: heigt,
	}
}

type Circle struct {
	Radius float64
}

// membuat nilai circle
func NewCircle(radius float64) Circle {
	return Circle{
		Radius: radius,
	}
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Parimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Parimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func Print(t string, s Shape) {
	fmt.Printf("%s area %v\n", t, s.Area())
	fmt.Printf("%s parimeter %v\n", t, s.Parimeter())
}

// disini volume bukan bagian dari interface shape
func (c Circle) Volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.Radius, 2)
}
